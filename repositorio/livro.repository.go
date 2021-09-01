package livro_repository

import (
	m "../models"
)

func Create(livro m.Livro) error {

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
