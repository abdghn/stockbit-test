/*
 * Created on 13/10/21 13.44
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abdghn/stockbit-test/microservice/api/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type server struct {
	proto.UnimplementedMovieSearchServer
}


func (s *server) HandleMovieSearch(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	searchword := req.GetSearchword()
	pagination := req.GetPagination()

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

	var responseObject proto.SearchResponse
	json.Unmarshal(responseData, &responseObject)


	return &responseObject, nil
}

func (s *server) DetailMovie(ctx context.Context, req *proto.SearchRequest) (*proto.Movie, error) {
	id := req.GetId()

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

	var responseObject proto.Movie
	json.Unmarshal(responseData, &responseObject)


	return &responseObject, nil
}



func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	gRPCServer := grpc.NewServer()
	proto.RegisterMovieSearchServer(gRPCServer,&server{})

	go func() {
		if err := gRPCServer.Serve(lis); err != nil {
			log.Fatalf("starting server failed: %v", err)
		}
	}()
	log.Printf("server is running at %s", ":3001")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server..")
	gRPCServer.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("gracefull shutdown success")
}

