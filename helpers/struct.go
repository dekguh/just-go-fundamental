package helpers

import "fmt"

type Engine struct {
	cc   int
	fuel string
}

type Vehicle struct {
	brand string
	model string
	Engine
}

func (v Vehicle) GetDetail() string {
	if v.brand == "" || v.model == "" || v.Engine.cc == 0 || v.Engine.fuel == "" {
		return "Vehicle is not valid"
	}

	return fmt.Sprintf("Brand: %s, Model: %s, Engine: %d cc, Fuel: %s", v.brand, v.model, v.Engine.cc, v.Engine.fuel)
}
