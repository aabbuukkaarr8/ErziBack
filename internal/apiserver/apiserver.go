package apiserver

import (
	"erzi_new/internal/handler/cart"
	userhalder "erzi_new/internal/handler/user"
	"net/http"
	"strings"

	"erzi_new/internal/handler/product"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
}

func New(config *Config) *APIServer {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return &APIServer{
		config: config,
		logger: logger,
		router: gin.Default(),
	}

}

func (s *APIServer) Run() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// конфигурация логгера
func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *APIServer) ConfigureRouter(prodHandler *product.Handler, cartHandler *cart.Handler, userHandler *userhalder.Handler) {
	s.router.POST("/user/create", userHandler.Create)
	s.router.POST("/user/login", userHandler.Login)

	s.router.GET("/products", prodHandler.GetAll)
	s.router.GET("/products/:id", prodHandler.GetByID)

	auth := s.router.Group("/auth", AuthMiddleware())
	{
		auth.POST("/products/create", RequireRole("admin"), prodHandler.Create)
		auth.PUT("/products/:id", RequireRole("admin"), prodHandler.Update)
		auth.DELETE("/products/:id", RequireRole("admin"), prodHandler.Delete)
		auth.POST("/cart/create", RequireRole("user", "admin"), cartHandler.CreateCart)
	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "роль не указана"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "роль имеет неверный формат"})
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		c.Abort()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
			c.Abort()
			return
		}
		var jwtSecret = []byte("your_secret_key")

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.HS256, jwtSecret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "некорректный токен"})
			c.Abort()
			return
		}

		userID, _ := token.Get("user_id")
		email, _ := token.Get("email")
		role, _ := token.Get("role")

		c.Set("user_id", userID)
		c.Set("email", email)
		c.Set("role", role)

		c.Next()
	}
}
