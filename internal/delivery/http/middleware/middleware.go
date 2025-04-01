package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/helper"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/auth"
)

type AuthMiddleware struct {
	cfg  *config.MainConfig
	repo repository.RepositoryImpl
}

func NewAuthMiddleware(cfg *config.MainConfig, repo repository.RepositoryImpl) *AuthMiddleware {
	return &AuthMiddleware{
		cfg:  cfg,
		repo: repo,
	}
}

func (m *AuthMiddleware) JWTAuth(h ...gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if authHeader == "" {
			helper.Error(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		barierToken := strings.Split(authHeader, "Bearer ")
		if len(barierToken) < 2 {
			helper.Error(c, http.StatusUnauthorized, "token not valid")
			return
		}

		token := barierToken[1]

		authJwt := auth.NewAuth(m.cfg)
		userID, err := authJwt.VerifyToken(token)
		if err != nil {
			helper.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := m.repo.GetUserRepo().GetByID(c, int(userID))
		if err != nil {
			helper.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		auth.SetUserContext(c, user)
		auth.SetTokenContext(c, token)

		if len(h) > 0 {
			h[0](c)
			return
		}

		c.Next()
	}
}
