package apiserver

import (
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
func (s *APIServer) ConfigureRouter(prodHandler *product.Handler) {
	s.router.POST("/products/create", prodHandler.Create)
	s.router.GET("/products/:id", prodHandler.GetByID)
	s.router.GET("/products", prodHandler.GetAll)
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
		role, _ := token.Get("role")

		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}
