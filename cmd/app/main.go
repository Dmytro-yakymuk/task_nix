package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/Dmytro-yakymuk/task_nix/models"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

const (
	url = "https://jsonplaceholder.typicode.com"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	posts := parsePosts()
	var wgPosts sync.WaitGroup

	for i := 0; i < len(posts); i++ {

		wgPosts.Add(1)
		go func(post *models.Post) {
			parseComments(post)
			defer wgPosts.Done()
		}(&posts[i])
	}

	wgPosts.Wait()
}

func parsePosts() []models.Post {
	resp, err := http.Get(url + "/posts?userId=7")
	if err != nil {
		log.Fatalf("Error GET request in address %s: %s", url+"/posts?userId=7", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	defer resp.Body.Close()

	var posts []models.Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatal(err.Error())
	}

	return posts
}

func parseComments(post *models.Post) {
	resp, err := http.Get(url + "/comments?postId=" + strconv.Itoa(post.Id))
	if err != nil {
		log.Fatalf("Error GET request in address %s: %s", url+"/comments?postId="+strconv.Itoa(post.Id), err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	defer resp.Body.Close()

	var comments []models.Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = savePosts(post); err != nil {
		log.Fatal(err.Error())
	}

	var wgComments sync.WaitGroup

	for i := 0; i < len(comments); i++ {
		wgComments.Add(1)
		go func(comment *models.Comment) {
			saveComments(comment)
			defer wgComments.Done()
		}(&comments[i])
	}
	wgComments.Wait()

}

func savePosts(post *models.Post) error {

	db := connectDB()
	defer db.Close()

	_, err := db.Exec(fmt.Sprintf("INSERT INTO `posts` (`user_id`, `id`, `title`, `body`) VALUE('%d', '%d', '%s', '%s')", post.UserId, post.Id, post.Title, post.Body))

	if err != nil {
		return err
	}

	return nil
}

func saveComments(comment *models.Comment) {

	db := connectDB()
	defer db.Close()

	_, err := db.Exec(fmt.Sprintf("INSERT INTO `comments` (`post_id`, `id`, `name`, `email`, `body`) VALUE('%d', '%d', '%s', '%s', '%s')", comment.PostId, comment.Id, comment.Name, comment.Email, comment.Body))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
