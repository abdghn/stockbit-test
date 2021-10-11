/*
 * Created on 11/10/21 14.51
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package handler

import (
	"encoding/json"
	"github.com/abdghn/stockbit-test/microservice/internal/usecase"
	"log"
	"net/http"
)

type Handler interface {
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
	HandleGetMovies(w http.ResponseWriter, r *http.Request)
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



func (h *handler) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	pokemon, err := h.uc.GetMovies()
	if err != nil {
		log.Printf("[handler.HandleGetMovies] unable to find user by id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	b, err := json.Marshal(&pokemon)
	if err != nil {
		log.Printf("[handler.HandleGetMovies] error while marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
