package main

import (
	"fmt"

	"gobootcamp.com/level1/change"
)

func main() {

	result, err := change.GiveChange(300, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
