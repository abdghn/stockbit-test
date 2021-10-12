/*
 * Created on 11/10/21 14.51
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/abdghn/stockbit-test/microservice/internal/model"
	"github.com/abdghn/stockbit-test/microservice/internal/resource/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Usecase interface {
	GetMovie(id string) (*model.Movie, error)
	GetMoviesSearch(pagination string, searchword string) (*model.Response, error)
}

type usecase struct {
	persistentDB db.Persistent
}

func New(persistentDB db.Persistent) Usecase {
	return &usecase{persistentDB: persistentDB}
}

func (u *usecase) GetMoviesSearch(pagination string, searchword string) (*model.Response, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", searchword, pagination)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)


	return &responseObject, nil

}

func (u *usecase) GetMovie(id string) (*model.Movie, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&i=%s", id)
	log.Printf(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Movie
	json.Unmarshal(responseData, &responseObject)


	return &responseObject, nil

}