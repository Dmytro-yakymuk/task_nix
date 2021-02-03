package main

import (
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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	url = "https://jsonplaceholder.typicode.com"
)

// Getenv takes values from environment variables, so you need to import them from a file .env
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

	savePosts(post)

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

func savePosts(post *models.Post) {
	db := connectDB()
	db.Create(post)
}

func saveComments(comment *models.Comment) {
	db := connectDB()
	db.Create(comment)
}

func connectDB() *gorm.DB {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func autoMigrate() {
	db := connectDB()
	db.AutoMigrate(&models.Post{}, &models.Comment{})
}
