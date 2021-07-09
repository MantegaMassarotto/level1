package main

import (
	"fmt"

	"gobootcamp.com/level1/queue"
)

func main() {
	//products := map[int]inventory.Product{2: {Id: 2, Name: "Product A"}}
	//invent := inventory.Inventory{Products: products}

	//newProduct := inventory.Product{}
	//newProduct.Id = 1
	//newProduct.Name = "Product A"

	//result, err := inventory.AddProduct(invent, newProduct)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//fmt.Println(result)

	fmt.Println(queue.Enqueue(1))
	fmt.Println(queue.Enqueue(2))
	fmt.Println(queue.Enqueue(3))
	fmt.Println(queue.Enqueue(4))
	fmt.Println(queue.Enqueue(5))
	fmt.Println(queue.Enqueue(6))

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
