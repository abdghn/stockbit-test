/*
 * Created on 11/10/21 14.51
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package handler

import (
	"encoding/json"
	"github.com/abdghn/stockbit-test/microservice/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler interface {
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
	HandleGetMovie(w http.ResponseWriter, r *http.Request)
	HandleSearchMovies(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc usecase.Usecase
}

// New creates a new handler
func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *handler) HandleSearchMovies(w http.ResponseWriter, r *http.Request) {
	pagination := r.URL.Query().Get("pagination")
	searchword := r.URL.Query().Get("searchword")
	movies, err := h.uc.GetMoviesSearch(pagination, searchword)
	if err != nil {
		log.Printf("[handler.HandleGetMovies] unable to find user by id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	b, err := json.Marshal(&movies)
	if err != nil {
		log.Printf("[handler.HandleGetMovies] error while marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}



func (h *handler) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	movie, err := h.uc.GetMovie(id)
	if err != nil {
		log.Printf("[handler.HandleGetMovie] unable to find user by id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	b, err := json.Marshal(&movie)
	if err != nil {
		log.Printf("[handler.HandleGetMovie] error while marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
