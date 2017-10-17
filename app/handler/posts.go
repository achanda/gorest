package handler

import (
	"encoding/json"
	"net/http"

	"github.com/achanda/gorest/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllPosts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	posts := []model.Post{}
	db.Find(&posts)
	respondJSON(w, http.StatusOK, posts)
}

func CreatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	post := model.Post{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&post).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, post)
}

func GetPost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	post := getPostOr404(db, title, w, r)
	if post == nil {
		return
	}
	respondJSON(w, http.StatusOK, post)
}

func UpdatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	post := getPostOr404(db, title, w, r)
	if post == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&post).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, post)
}

func DeletePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	post := getPostOr404(db, title, w, r)
	if post == nil {
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getPostOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Post {
	post := model.Post{}
	if err := db.First(&post, model.Post{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &post
}
