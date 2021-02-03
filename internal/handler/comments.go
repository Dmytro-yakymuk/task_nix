package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllComments(c echo.Context) error {
	var comments []models.Comment

	comments, err := h.services.Comments.GetAll()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	accept := fmt.Sprintf("%v", c.Request().Header["Accept"])
	c.Response().Header().Set("Content-Type", accept)

	switch accept {
	case "[application/xml]":
		return c.XML(http.StatusOK, comments)
	case "[application/json]":
		return c.JSON(http.StatusOK, comments)
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

func (h *Handler) createComment(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = h.services.Comments.Create(&comment)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) getOneComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	comment, err := h.services.Comments.GetOne(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	accept := fmt.Sprintf("%v", c.Request().Header["Accept"])
	c.Response().Header().Set("Content-Type", accept)

	switch accept {
	case "[application/xml]":
		return c.XML(http.StatusOK, comment)
	case "[application/json]":
		return c.JSON(http.StatusOK, comment)
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

func (h *Handler) updateComment(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var comment models.Comment
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = json.Unmarshal(body, &comment)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	comment.Id = id

	err = h.services.Comments.Update(&comment)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) deleteComment(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = h.services.Comments.Delete(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)

}
