package controllers

import (
	"biblioteca/models"
	"encoding/json"
	"log"
	"net/http"
)

func Index(write http.ResponseWriter, read *http.Request) {
	livros, err := models.SearchLivro()
	if err != nil {
		http.Error(write, "Houve um erro ao encontrar os livros", http.StatusInternalServerError)
		log.Println("Erro ao encontrar os livros: ", err)
		return
	}
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(livros)
}

func Create(write http.ResponseWriter, read *http.Request) {
	if read.Method != "POST" {
		http.Error(write, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	var livro models.Livro
	err := json.NewDecoder(read.Body).Decode(&livro)
	if err != nil {
		http.Error(write, "Erro ao decodificar o livro", http.StatusBadRequest)
		return
	}

	err = models.CreateLivro(livro)
	if err != nil {
		log.Println("Erro ao inserir o livro: ", err)
		http.Error(write, "Erro ao inserir o livro", http.StatusInternalServerError)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(livro)
}

func Delete(write http.ResponseWriter, read *http.Request) {
	if read.Method != "DELETE" {
		http.Error(write, "Metodo invalido", http.StatusMethodNotAllowed)
		return
	}

	idLivro := read.URL.Query().Get("id")
	err := models.DeleteLivro(idLivro)
	if err != nil {
		log.Println("Erro ao deletar o livro: ", err)
		http.Error(write, "Erro ao deletar o livro", http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusOK)
}
