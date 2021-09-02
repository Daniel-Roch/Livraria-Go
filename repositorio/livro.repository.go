package livro_repository

import (
	"context"

	"../database"
	m "../models"
)

var collection = database.GetCollection("livros")
var ctx = context.Background()

func Create(livro m.Livro) error {

	var err error

	_, err = collection.InsertOne(ctx, livro)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Livros, error) {

	var livros m.Livros

	return livros, nil
}

func Update(livro m.Livro, livroId int) error {

	return nil
}

func Delete(livrId int) error {

	return nil
}
