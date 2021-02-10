package handler

import (
	"net/http"

	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// auth godoc
// @Summary auth user
// @Tags auth
// @Description auth user
// @ID auth
// @Accept json
// @Produce json
// @Success 200 {int} echo.Context.Response().Status
// @Router /auth [get]
func (h *Handler) auth(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return c.NoContent(http.StatusOK)
}

// authCallback godoc
// @Summary auth callback user
// @Tags auth
// @Description auth and create user
// @ID authCallback
// @Accept json
// @Produce json
// @Success 200 {int} echo.Context.Response().Status
// @Success 201 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /auth/callback [get]
func (h *Handler) authCallback(c echo.Context) error {

	value, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	userID, err := gothic.GetFromSession("userID", c.Request())
	if err != nil && userID != "" {
		return c.NoContent(http.StatusNotFound)
	}
	if userID != value.UserID || userID == "" {
		err = gothic.StoreInSession("userID", value.UserID, c.Request(), c.Response())
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		user := models.User{Id: value.UserID, Name: value.Name, Email: value.Email}

		err = h.services.Users.Create(&user)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusCreated)
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// logout godoc
// @Summary logout user
// @Tags auth
// @Description logout and delete user
// @ID logout
// @Accept json
// @Produce json
// @Success 204 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /logout [get]
func (h *Handler) logout(c echo.Context) error {
	id, err := gothic.GetFromSession("userID", c.Request())
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	gothic.Logout(c.Response(), c.Request())
	err = h.services.Users.Delete(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)
}
