package user

import (
	"net/http"
	"url-shortening-service/pkg/domain"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register
// @Description Registra un usuario.
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body domain.RegisterInput true "Datos de registro"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func (h *Handler) Register(ctx *gin.Context) {
	var input domain.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPwsd, err := h.Service.EncryptPassword(input.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "errot to try encrypt pwsd",
		})
		return
	}

	user := domain.RegisterInput{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPwsd,
	}

	err = h.Service.PostUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error to try create user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"username": input.Username,
		"email":    input.Email,
		"password": input.Password,
	})

}
