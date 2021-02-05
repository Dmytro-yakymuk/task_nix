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

// getAllComments godoc
// @Summary List comments
// @Tags comments
// @Description get all comments
// @ID getAllComments
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {array} models.Comment
// @Failure 404 {int} echo.Context.Response().Status
// @Router /comments [get]
func (h *Handler) getAllComments(c echo.Context) error {
	var comments []models.Comment

	comments, err := h.services.Comments.GetAll()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	accept := fmt.Sprintf("%v", c.Request().Header["Accept"])
	c.Response().Header().Set("Content-Type", accept)

	switch accept {
	case "[text/xml]":
		return c.XML(http.StatusOK, comments)
	case "[application/json]":
		return c.JSON(http.StatusOK, comments)
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

// createComment godoc
// @Tags comments
// @Summary Add a comment
// @Description create comment
// @ID createComment
// @Accept json
// @Produce json
// @Param input body models.Comment true "info for comment"
// @Success 201 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /comments [post]
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

// getOneComment godoc
// @Summary Show a comment
// @Tags comments
// @Description show comment for input id
// @ID getOneComment
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Param id path string true "id for comment"
// @Success 200 {object} models.Comment
// @Failure 404 {int} echo.Context.Response().Status
// @Router /comments/{id} [get]
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
	case "[text/xml]":
		return c.XML(http.StatusOK, comment)
	case "[application/json]":
		return c.JSON(http.StatusOK, comment)
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

// updateComment godoc
// @Tags comments
// @Summary Update a comment
// @Description update comment
// @ID updateComment
// @Accept json
// @Produce json
// @Param id path string true "id for comment"
// @Param input body models.Comment true "info for comment"
// @Success 204 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /comments/{id} [put]
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

// deleteComment godoc
// @Tags comments
// @Summary Delete a comment
// @Description delete comment
// @ID deletecomment
// @Accept json
// @Produce json
// @Param id path string true "id for comment"
// @Success 204 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /comments/{id} [delete]
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
