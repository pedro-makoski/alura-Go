package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	password := os.Getenv("PASSWORD_DB")

	conexao := fmt.Sprintf("user=postgres dbname=alura_loja password=%s host=localhost sslmode=disable", password)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}