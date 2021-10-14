/*
 * Created on 11/10/21 16.07
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package model

// A Response struct to map the Entire Response
type HTTPResponse struct {
	Result      []Movie `json:"Search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchRequest struct {
	Searchword string
	Pagination      string
}

type SearchResponse struct {
	Response     []Movie `json:"Search"`
	ErrorMessage string  `json:"Error,omitempty"`
}
