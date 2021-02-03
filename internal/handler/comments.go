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
	urlComments   = "/api/v1/comments"
	urlCommentsId = regexp.MustCompile(`/api/v1/comments/+\d+$`)
)

func (h *Handler) routeComments(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		if r.URL.Path == urlComments {
			h.getAllComments(w, r)
		} else if urlCommentsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.getOneComment(w, r, id)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "POST":
		if r.URL.Path == urlComments {
			h.createComment(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "PUT":
		if urlCommentsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.updateComment(w, r, id)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case "DELETE":
		if urlCommentsId.MatchString(r.URL.Path) {
			path := strings.Split(r.URL.Path, "/")
			id, err := strconv.Atoi(path[4])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				h.deleteComment(w, r, id)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (h *Handler) getAllComments(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment
	accept := fmt.Sprintf("%v", r.Header["Accept"])

	w.Header().Set("Content-Type", accept)

	comments, err := h.services.Comments.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		switch accept {
		case "[application/xml]":

			res, err := xml.Marshal(comments)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))

		case "[application/json]":
			res, err := json.Marshal(comments)
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

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	err = h.services.Comments.Create(&comment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getOneComment(w http.ResponseWriter, r *http.Request, id int) {
	comment, err := h.services.Comments.GetOne(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		accept := fmt.Sprintf("%v", r.Header["Accept"])
		w.Header().Set("Content-Type", accept)

		switch accept {
		case "[application/xml]":

			res, err := xml.Marshal(comment)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))

		case "[application/json]":
			res, err := json.Marshal(comment)
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

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request, id int) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	comment.Id = id

	err = h.services.Comments.Update(&comment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request, id int) {

	err := h.services.Comments.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)

}
