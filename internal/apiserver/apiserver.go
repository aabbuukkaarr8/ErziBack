package apiserver

import (
	"erzi_new/internal/handler/cart"
	"erzi_new/internal/handler/cartItem"
	userhandler "erzi_new/internal/handler/user"
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

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *APIServer) ConfigureRouter(prodHandler *product.Handler, userHandler *userhandler.Handler, cartitemHandler *cartItem.Handler, cartHandler *cart.Handler) {
	s.router.POST("/user/register", userHandler.Create)
	s.router.POST("/user/login", userHandler.Login)

	s.router.GET("/products", prodHandler.GetAll)
	s.router.GET("/products/:id", prodHandler.GetByID)
	protected := s.router.Group("/", AuthMiddleware())

	{
		protected.POST("/cart/delete_all", cartitemHandler.DeleteAll)
		protected.DELETE("cart/items/:id/delete", cartitemHandler.DeleteCartItem)
		protected.PUT("/cart/items/:id/increment", cartitemHandler.IncrementQuantity)
		protected.PUT("/cart/items/:id/decrement", cartitemHandler.DecrementQuantity)
		protected.POST("/:product_id/add_to_cart", cartitemHandler.AddCartItem)
		protected.GET("/cart/items", cartitemHandler.GetAllCartItems)
		protected.POST("/attribute/create", RequireRole("admin"), prodHandler.CreateAttribute)
		protected.POST("/products/create", RequireRole("admin"), prodHandler.Create)
		protected.PUT("/products/:id", RequireRole("admin"), prodHandler.Update)
		protected.DELETE("/products/:id", RequireRole("admin"), prodHandler.Delete)
		protected.GET("/cart/restore", cartHandler.Restore)
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

		userID, _ := token.Get("userID")
		email, _ := token.Get("email")
		role, _ := token.Get("role")

		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("role", role)

		c.Next()
	}
}
