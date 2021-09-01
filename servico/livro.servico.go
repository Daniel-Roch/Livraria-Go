package livro_servico

import (
	m "../models"
	livroRepository "../repositorio"
)

func Create(livro m.Livro) error {

	err := livroRepository.Create(livro)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Livros, error) {

	livros, err := livroRepository.Read()

	if err != nil {
		return nil, err
	}

	return livros, nil
}

func Update(livro m.Livro, livroId int) error {

	err := livroRepository.Update(livro, livroId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(livroId int) error {

	err := livroRepository.Delete(livroId)

	if err != nil {
		return err
	}

	return nil
}
