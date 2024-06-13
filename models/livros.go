package models

import (
	"biblioteca/database"
	"database/sql"
)

var DB *sql.DB

type Livro struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Autor      string  `json:"autor"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

func CreateLivro(l Livro) error {
	stmt, err := database.DB.Prepare("INSERT INTO livros (nome, autor, quantidade, preco) VALUES($1,$2,$3,$4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(l.Nome, l.Autor, l.Quantidade, l.Preco)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func SearchLivro() ([]Livro, error) {
	allLivros, err := database.DB.Query("SELECT * FROM livros ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	livros := []Livro{}

	for allLivros.Next() {
		var l Livro
		err = allLivros.Scan(&l.Id, &l.Nome, &l.Autor, &l.Quantidade, &l.Preco)
		if err != nil { // checa por erro no scan antes de dar append
			return nil, err
		}
		livros = append(livros, l)
	}

	defer allLivros.Close()

	return livros, nil
}

func UpdateLivro(id int, nome, autor string, quantidade int, preco float64) error {
	stmt, err := database.DB.Prepare("UPDATE livros SET nome = $1, autor = $2, quantidade = $3, preco = $4 WHERE id = $5")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nome, autor, quantidade, preco, id)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteLivro(id string) error {
	stmt, err := database.DB.Prepare("DELETE FROM livros WHERE id = $1") // stmt vulgo statement eh a query preparada
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
