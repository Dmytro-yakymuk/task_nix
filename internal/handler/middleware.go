package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, err := gothic.GetFromSession("userID", c.Request())
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		user, err := h.services.Users.GetOne(userID)
		if err != nil || user == nil {
			return c.NoContent(http.StatusNotFound)
		}

		return next(c)
	}
}
