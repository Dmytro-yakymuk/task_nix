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

func (h *Handler) getAllPosts(c echo.Context) error {
	var posts []models.Post
	accept := fmt.Sprintf("%v", c.Request().Header["Accept"])
	c.Response().Header().Set("Content-Type", accept)

	posts, err := h.services.Posts.GetAll()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	switch accept {
	case "[application/xml]":
		return c.XML(http.StatusOK, posts)
	case "[application/json]":
		return c.JSON(http.StatusOK, posts)
	default:
		return c.NoContent(http.StatusNotFound)
	}

}

func (h *Handler) createPost(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = h.services.Posts.Create(&post)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) getOnePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	post, err := h.services.Posts.GetOne(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	accept := fmt.Sprintf("%v", c.Request().Header["Accept"])
	c.Response().Header().Set("Content-Type", accept)

	switch accept {
	case "[application/xml]":
		return c.XML(http.StatusOK, post)
	case "[application/json]":
		return c.JSON(http.StatusOK, post)
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

func (h *Handler) updatePost(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var post models.Post
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	post.Id = id

	err = h.services.Posts.Update(&post)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) deletePost(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	err = h.services.Posts.Delete(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusNoContent)

}
