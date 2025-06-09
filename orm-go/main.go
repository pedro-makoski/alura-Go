package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome      string
	Area      string
	Professor string
}

func main() {
	db, err := gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao acessar o banco de dados")
	}
	db.AutoMigrate(&Curso{})
	db.Create(&Curso{Nome: "HTML, CSS", Area: "Desenvolvimento WEB", Professor: "Gustavo Guanabara"})
	var curso Curso
	db.First(&curso, 1)
	fmt.Println(curso)
	db.Model(&curso).Update("Nome", "HTML 5 e CSS 3")
	db.First(&curso, 1)
	fmt.Println(curso)
	db.Delete(&curso, 1)
}
