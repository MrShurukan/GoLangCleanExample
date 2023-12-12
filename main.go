package main

import (
	"fmt"
	"testClean/application"
)

func main() {
	fmt.Println("Hello! Here are goods:")

	goods, err := application.GetAllGoods()
	if err != nil {
		fmt.Println("Возникла ошибка при получении:")
		fmt.Println(err)
	}

	for i, good := range goods {
		fmt.Printf("%d) %s (%f)\n", i+1, good.Name, good.Price)
	}
}
