package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/achanda/gorest/app/handler"
	"github.com/achanda/gorest/app/model"
	"github.com/achanda/gorest/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/posts", a.GetAllPosts)
	a.Post("/posts", a.CreatePost)
	a.Get("/posts/{title}", a.GetPost)
	a.Delete("/posts/{title}", a.DeletePost)

	a.Get("/version", a.GetVersion)
}

func (a *App) GetVersion(w http.ResponseWriter, r *http.Request) {
	handler.GetVersion(a.DB, w, r)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPosts(a.DB, w, r)
}

func (a *App) CreatePost(w http.ResponseWriter, r *http.Request) {
	handler.CreatePost(a.DB, w, r)
}

func (a *App) GetPost(w http.ResponseWriter, r *http.Request) {
	handler.GetPost(a.DB, w, r)
}

func (a *App) DeletePost(w http.ResponseWriter, r *http.Request) {
	handler.DeletePost(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
