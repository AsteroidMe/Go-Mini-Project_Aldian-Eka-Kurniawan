package controller

import (
	"eco-journal/entities"
	"eco-journal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Failed input data: " + err.Error(),
			},
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Fill email and password",
			},
		})
		return
	}

	createdUser, err := uc.userService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Register failed: " + err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, entities.Response{
		Meta: entities.Meta{
			Status:  true,
			Message: "Register success",
		},
		Data: createdUser,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Failed input data: " + err.Error(),
			},
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Fill email and password",
			},
		})
		return
	}

	token, err := uc.userService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entities.Response{
			Meta: entities.Meta{
				Status:  false,
				Message: "Invalid credentials: " + err.Error(),
			},
		})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, entities.Response{
		Meta: entities.Meta{
			Status:  true,
			Message: "Login success",
		},
		Data: map[string]string{"token": token},
	})
}
