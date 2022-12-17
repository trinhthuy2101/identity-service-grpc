// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"

	"ecommerce/identity/internal/usecase"
)

func RegisterRouter(handler *gin.RouterGroup, t usecase.AuthUsecase) {
	newIdentityRoutes(handler, t)
}
