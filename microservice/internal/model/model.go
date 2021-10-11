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
