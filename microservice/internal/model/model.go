/*
 * Created on 11/10/21 16.07
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package model

// A Response struct to map the Entire Response
type Response struct {
	Search []Search `json:"search"`
}

type Search struct {
	Title string `json:"title"`
	Year string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"type"`
	Poster string `json:"poster"`
}

type Movie struct {
	Title string `json:"title"`
	Year string `json:"year"`
	Rated string `json:"rated"`
	Released string `json:"released"`
	Runtime string `json:"runtime"`
	Genre string `json:"genre"`
	Director string `json:"director"`
	Writer string `json:"writer"`
	Actors string `json:"actors"`
	Plot string `json:"plot"`
	Language string `json:"language"`
	Country string `json:"country"`
	Awards string `json:"awards"`
	Poster string `json:"poster"`
	Ratings []Rating `json:"ratings"`
	Metascore string `json:"metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes string `json:"imdbVotes"`
	ImdbID string `json:"imdbID"`
	Type string `json:"type"`
	DVD string `json:"dvd"`
	BoxOffice string `json:"boxOffice"`
	Production string `json:"production"`
	Website string `json:"website"`
	Response string `json:"response"`
}

type Rating struct {
	Source string `json:"source"`
	Value string `json:"value"`
}
