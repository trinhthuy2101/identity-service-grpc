package v1

import (
	"github.com/gin-gonic/gin"

	"ecommerce/identity/internal/entity"
	"ecommerce/identity/internal/usecase"
	"ecommerce/identity/pkg/errors"
	"ecommerce/identity/pkg/logger"
	"ecommerce/identity/pkg/response"
)

type identityRoutes struct {
	u usecase.AuthUsecase
}

func newIdentityRoutes(handler *gin.RouterGroup, t usecase.AuthUsecase) {
	r := &identityRoutes{t}

	h := handler.Group("/auth")
	{
		h.GET("/", response.GinWrap(r.login))
		h.GET("/:id", response.GinWrap(r.register))
	}
}

// @Summary     login
// @Description login
// @Tags  	    Identity
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=response.Collection{results=[]entity.AdminUser}}
// @Failure     500 {object} response.FailureResponse
// @Router      /v1/login [post]
func (r *identityRoutes) login(c *gin.Context) *response.Response {
	var request AuthenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		return response.BadRequest()
	}

	adminUser, token, err := r.u.Login(c.Request.Context(), &entity.AdminUser{Email: request.Email, Password: request.Password})
	if err != nil {
		logger.Error("failed to login. Error %w", err)

		return response.Failure(err)
	}

	return response.SuccessWithData(
		&loginResponse{
			ID:       adminUser.ID,
			Email:    adminUser.Email,
			Token:    token,
			FullName: adminUser.FullName,
			Address:  adminUser.Address,
		},
	)
}

// @Summary     Register
// @Description Regisetr
// @Tags  	    Identity
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=response.Data{result=response.Success()}}
// @Failure     500 {object} response.FailureResponse
// @Router      /v1/register [post]
func (r *identityRoutes) register(c *gin.Context) *response.Response {
	var request AuthenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		return response.BadRequest()
	}

	err := r.u.Register(c.Request.Context(), &entity.AdminUser{Email: request.Email, Password: request.Password})
	if err != nil {
		logger.Error(err)
		err = errors.ErrCustomerNotFound

		return response.Failure(err)
	}

	return response.Success()
}
