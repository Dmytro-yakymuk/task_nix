package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dmytro-yakymuk/task_nix/internal/models"
)

var (
	urlPosts   = "/api/v1/posts"
	urlPostsId = regexp.MustCompile(`/api/v1/posts/+\d+$`)
)

func (h *Handler) routePosts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		if r.URL.Path == urlPosts {
			h.getAllPosts(w, r)
		} else if urlPostsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.getOnePost(w, r, id)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "POST":
		if r.URL.Path == urlPosts {
			h.createPost(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "PUT":
		if urlPostsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.updatePost(w, r, id)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "DELETE":
		if urlPostsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.deletePost(w, r, id)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	accept := fmt.Sprintf("%v", r.Header["Accept"])

	w.Header().Set("Content-Type", accept)

	posts, err := h.services.Posts.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		switch accept {
		case "[application/xml]":

			res, err := xml.Marshal(posts)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))

		case "[application/json]":
			res, err := json.Marshal(posts)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	err = h.services.Posts.Create(&post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getOnePost(w http.ResponseWriter, r *http.Request, id int) {
	post, err := h.services.Posts.GetOne(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		accept := fmt.Sprintf("%v", r.Header["Accept"])
		w.Header().Set("Content-Type", accept)

		switch accept {
		case "[application/xml]":

			res, err := xml.Marshal(post)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))

		case "[application/json]":
			res, err := json.Marshal(post)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request, id int) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	post.Id = id

	err = h.services.Posts.Update(&post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request, id int) {

	err := h.services.Posts.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)

}
