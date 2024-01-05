package handler

import (
	"dating-apps/domain/users"
	"dating-apps/domain/users/model"
	"dating-apps/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	usecase users.UsercaseInterface
}

func NewHandler(usecase users.UsercaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
	}
}

func (h *Handler) Register(c *gin.Context) {
	var payload model.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Bind JSON",
			"error":   err.Error(),
		})
		return
	}

	err = h.usecase.Register(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Register",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (h *Handler) Login(c *gin.Context) {
	var payload model.UserLogin

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Bind JSON",
			"error":   err.Error(),
		})
		return
	}

	token, err := h.usecase.Login(payload)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   token,
	})
}

func (h *Handler) GetProfiles(c *gin.Context) {
	// Get users information from token
	data := c.MustGet("userData").(jwt.MapClaims)
	email := data["email"].(string)

	users, err := h.usecase.GetProfiles(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Get Profiles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    users,
	})
}

func (h *Handler) Swipe(c *gin.Context) {
	data := c.MustGet("userData").(jwt.MapClaims)
	email := data["email"].(string)
	subscription := data["subscription"].(string)

	id := c.Param("id")
	action := c.Param("action")

	err := h.usecase.Swipe(email, id, action, subscription)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Swipe",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (h *Handler) Purchase(c *gin.Context) {
	data := c.MustGet("userData").(jwt.MapClaims)
	email := data["email"].(string)

	err := h.usecase.Purchase(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Purchase",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
