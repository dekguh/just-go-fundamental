package cases

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID          string
	Description string
	Amount      float64
	Status      string
}

func MakeOrder(order Order, wg *sync.WaitGroup, response chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	response <- fmt.Sprintf("Order %s is processed", order.ID)
}

func InvoiceOrder(order Order, wg *sync.WaitGroup, response chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	response <- fmt.Sprintf("Invoice %s with total amount %f", order.ID, order.Amount)
}

func OrderBouquets() {
	orders := []Order{
		{ID: "1", Description: "Bouquet 1", Amount: 100000},
		{ID: "2", Description: "Bouquet 2", Amount: 200000},
		{ID: "3", Description: "Bouquet 3", Amount: 300000},
	}
	response := make(chan string, len(orders)*2)
	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(2)
		go MakeOrder(order, &wg, response)
		go InvoiceOrder(order, &wg, response)
	}

	go func() {
		wg.Wait()
		close(response)
	}()

	for result := range response {
		fmt.Println(result)
	}

	fmt.Println("All orders are processed")
}
