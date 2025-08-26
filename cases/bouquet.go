package cases

type Shop interface {
	GetName() string
	GetAddress() string
}

type FlowerShop struct {
	Name    string
	Address string
}

func (f *FlowerShop) GetName() string {
	return f.Name
}

func (f *FlowerShop) GetAddress() string {
	return f.Address
}
