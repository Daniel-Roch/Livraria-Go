package models

import "time"

//Dados de Livro
type Livro struct {
	ID        int       `json:"_id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Isnb10    string    `json:"isnb-10"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

//Lista de Livros
type Livros []*Livro
