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

// getAllPosts godoc
// @Summary List posts
// @Tags posts
// @Description get all posts
// @ID getAllPosts
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {array} models.Post
// @Failure 404 {int} echo.Context.Response().Status
// @Router /posts [get]
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

// createPost godoc
// @Tags posts
// @Summary Add a post
// @Description create post
// @ID createPost
// @Accept json
// @Produce json
// @Param input body models.Post true "info for post"
// @Success 201 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /posts [post]
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

// getOnePost godoc
// @Summary Show a post
// @Tags posts
// @Description show post for input id
// @ID getOnePost
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Param id path string true "id for post"
// @Success 200 {object} models.Post
// @Failure 404 {int} echo.Context.Response().Status
// @Router /posts/{id} [get]
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

// updatePost godoc
// @Tags posts
// @Summary Update a post
// @Description update post
// @ID updatePost
// @Accept json
// @Produce json
// @Param id path string true "id for post"
// @Param input body models.Post true "info for post"
// @Success 204 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /posts/{id} [put]
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

// deletePost godoc
// @Tags posts
// @Summary Delete a post
// @Description delete post
// @ID deletePost
// @Accept json
// @Produce json
// @Param id path string true "id for post"
// @Success 204 {int} echo.Context.Response().Status
// @Failure 404 {int} echo.Context.Response().Status
// @Router /posts/{id} [delete]
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
