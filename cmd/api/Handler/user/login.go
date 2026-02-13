package user

import (
	"net/http"
	"strings"
	"url-shortening-service/pkg/domain"
	"url-shortening-service/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login
// @Description Autentica un usuario y retorna un token JWT (también se setea cookie `access_token`).
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body domain.LoginInput true "Credenciales"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func (h *Handler) Login(ctx *gin.Context) {

	var user domain.LoginInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	user.Email = strings.TrimSpace(user.Email)

	u, err := h.Service.GetUser(domain.LoginInput{Email: user.Email})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user don't exist",
		})
		return
	}

	if u == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user don't exist",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := auth.GenerateToken(user.Email, "")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	ctx.SetCookie(
		"access_token",
		token.Token,
		3600,
		"/",
		"",
		ctx.Request.TLS != nil,
		true,
	)

	ctx.Set("email", token.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
