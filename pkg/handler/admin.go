package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/GO-PROJECT/pkg/usecases"
	"github.com/sgokul961/GO-PROJECT/pkg/utils/models"
	"github.com/sgokul961/GO-PROJECT/pkg/utils/response"
)

type AdminHandler struct {
	adminUseCase usecases.AdminUseCaseinterface
}

func NewAdminHandler(usecase usecases.AdminUseCaseinterface) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin

	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}
