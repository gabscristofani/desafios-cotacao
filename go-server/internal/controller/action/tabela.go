package controller

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CriarTabela() {
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, valor TEXT, data TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabela criada com sucesso")

}
