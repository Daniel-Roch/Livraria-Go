package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Author    string             `json:"author"`
	Isnb10    string             `json:"isnb-10"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

var client *mongo.Client

func CreateBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	collection := client.Database("Livraria_Moneri").Collection("Livro")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, book)
	json.NewEncoder(response).Encode(result)
}

//Fazer <-
func GetBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var books []Book
	collection := client.Database("Livraria_Moneri").Collection("Livro")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"mensagem": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"mensagem": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(books)
}

func main() {
	fmt.Printf("Aplicação Rodando..")
	//Decidi não pegar o erro no momento.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27016")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book", CreateBooks).Methods("POST")
	http.ListenAndServe(":8000", router)
}
