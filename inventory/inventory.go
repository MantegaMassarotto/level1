package inventory

import "errors"

type Product struct {
	Id   int
	Name string
}

type Inventory struct {
	Products map[int]Product
}

func AddProduct(inventory Inventory, product Product) (Inventory, error) {
	_, ok := inventory.Products[product.Id]
	if ok {
		return inventory, errors.New("Product already exits in inventory")
	}
	inventory.Products[product.Id] = product
	return inventory, nil
}
