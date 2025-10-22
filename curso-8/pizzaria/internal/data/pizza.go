package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
)

var pizzas = []models.Pizza{}

func LoadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de pizzas.json")
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Erro ao dessirealizar o JSON")
		return
	}
}

func SavePizza() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de pizzas.json")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ")
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON: ", err.Error())
	}
}

func GetPizzas() []models.Pizza {
	return pizzas
}

func AddPiza(newPizza models.Pizza) error {
	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		return err
	}

	pizzas = append(pizzas, newPizza)
	SavePizza()

	return nil
}

func RemovPizza(idx int) {
	pizzas = append(pizzas[:idx], pizzas[idx+1:]...)
	SavePizza()
}

func UpdatePizza(idx int, newPizza models.Pizza, id int) error {
	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		return err
	}

	for _, review := range newPizza.Reviews {
		if err := service.ValidateRatingReview(review); err != nil {
			return err
		}
	}

	pizzas[idx] = newPizza
	pizzas[idx].ID = id
	SavePizza()

	return nil
}
