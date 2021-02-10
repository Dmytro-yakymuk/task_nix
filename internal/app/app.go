package app

import (
	"log"
	"os"

	"github.com/Dmytro-yakymuk/task_nix/internal/config"
	"github.com/Dmytro-yakymuk/task_nix/internal/handler"
	"github.com/Dmytro-yakymuk/task_nix/internal/models"
	"github.com/Dmytro-yakymuk/task_nix/internal/repository"
	"github.com/Dmytro-yakymuk/task_nix/internal/service"
	"github.com/Dmytro-yakymuk/task_nix/pkg/database/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db := mysql.ConnectDB(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Name)
	e := echo.New()

	// Migrations
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// OAuth2.0
	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:8080/auth/callback?provider=twitter"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),

		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:8080/auth/callback?provider=facebook", "email"),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/callback?provider=google", "email", "profile"),
	)

	// Services, Repos & API Handlers
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	handlers.Init(e)

	// HTTP Server
	e.Logger.Fatal(e.Start(":" + cfg.HTTP.Port))

}

/*
http://localhost:8080/auth?provider=google
*/
