package main

import (
	"fmt"
	"learn/just-go-fundamental/cases"
	"learn/just-go-fundamental/helpers"
)

func Hello(name string) string {
	return name
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helpers.TextCapitalize("hello world"))
	fmt.Println(helpers.InterpolateFormula(1, 10, 20))
	fmt.Println(helpers.GenerateCarDetail())

	var balance int = 0
	var wallet *int = &balance
	*wallet = 20
	fmt.Println(balance)

	totalFruit := map[string]int{
		"apple":      10,
		"banana":     20,
		"cherry":     30,
		"date":       40,
		"elderberry": 50,
	}
	fmt.Println(totalFruit)
	totalFruit["watermelon"] = 999
	fmt.Println(totalFruit)

	type User struct {
		ID    int
		name  string
		email string
	}

	var users = make(map[int]User)
	users[1] = User{
		ID:    1,
		name:  "John Doe",
		email: "john.doe@example.com",
	}

	for _, user := range users {
		fmt.Println(user)
	}

	var flowerShopDetail cases.Shop
	gorgeousFleur := cases.FlowerShop{
		Name:    "Gorgeous Fleur",
		Address: "Jl. Raya Cibadak No. 123",
	}
	flowerShopDetail = &gorgeousFleur

	fmt.Println(flowerShopDetail.GetName())
	fmt.Println(flowerShopDetail.GetAddress())

	gorgeousFleur.Address = "JL By Pass Ngurah Rai"
	fmt.Println(flowerShopDetail.GetAddress())

	type Sum[T any] struct {
		Total T
	}
	triangle := Sum[float64]{
		Total: 12.5,
	}
	fmt.Println(triangle.Total)
}
