package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Dmytro-yakymuk/task_nix/internal/models"
)

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		var posts []models.Post

		posts, err := h.services.Posts.GetAll()
		if err != nil {
			log.Fatal(err.Error())
		}

		keys, ok := r.URL.Query()["type"]

		if ok && keys[0] == "xml" {
			w.Header().Set("Content-Type", "application/xml")
			xml, err := xml.MarshalIndent(posts, " ", "  ")
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(xml))

		} else {

			w.Header().Set("Content-Type", "application/json")
			json, err := json.Marshal(posts)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(json))
		}

	case "POST":

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
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}
