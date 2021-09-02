package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Author    string             `json:"author"`
	Isnb10    string             `json:"isnb-10"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}

var client *mongo.Client

func CreateBooksEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	collection := client.Database("Livraria_Moneri").Collection("Livro")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, book)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Printf("Aplicação Rodando..")
	//Decidi não pegar o erro no momento.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27016")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/book", CreateBooksEndpoint).Methods("POST")
	http.ListenAndServe(":8000", router)
}
