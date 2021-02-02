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
			h.getAll(w, r)
		} else if urlPostsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.getOne(w, r, id)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "POST":
		if r.URL.Path == urlPosts {
			h.Create(w, r)
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
				h.Update(w, r, id)
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
				h.Delete(w, r, id)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	accept := fmt.Sprintf("%v", r.Header["Accept"])

	w.Header().Set("Content-Type", accept)

	posts, err := h.services.Posts.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf(accept)

	if accept == "[application/xml]" {
		res, err := xml.MarshalIndent(posts, " ", "  ")
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))

	} else {
		res, err := json.Marshal(posts)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
	}

	err = h.services.Posts.Create(&post)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	post, err := h.services.Posts.GetOne(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		keys, ok := r.URL.Query()["type"]

		if ok && keys[0] == "xml" {
			w.Header().Set("Content-Type", "application/xml")
			xml, err := xml.MarshalIndent(post, " ", "  ")
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(xml))

		} else {

			w.Header().Set("Content-Type", "application/json")
			json, err := json.Marshal(post)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(json))
		}
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, id int) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
	}
	post.Id = id

	err = h.services.Posts.Update(&post)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, id int) {

	err := h.services.Posts.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
