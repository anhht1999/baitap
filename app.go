package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"strconv"
    // "strings"
	// "sort"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
func main() {

	//in ra mảng sản phầm
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}
	for _, product := range products {
		fmt.Println(product)
	}

	// fmt.Println("------Sắp xếp theo giá cao nhất--------")
	// sortedProducts := products[:]
	// sort.Slice(sortedProducts, func(i, j int) bool {
	// 	return sortedProducts[i].Price > sortedProducts[j].Price
	// })
	// for _, product := range sortedProducts {
	// 	fmt.Println(product)
	// }

	// fmt.Println("------Top 5 giá cao nhất--------")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(sortedProducts[i])
	// }

	// fmt.Println("------Đếm sản phẩm theo category--------")
	// myMap := make(map[string]int)
	// for _, product := range products {
	// 	myMap[product.Category]++
	// }
	// fmt.Println(myMap)

	fmt.Println("------Loc Product--------")
	fmt.Println("Nhap vao ten product")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	for _, product := range products {
		if product.Name == input {
			fmt.Println(product)
			break
		}
	}

	fmt.Println("------Loc Category--------")

	fmt.Println("Nhap vao ten category")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	input1 := scanner1.Text()

	for _, product := range products {
		if product.Category == input1 {
			fmt.Println(product)
		}
	}

	fmt.Println("------Loc tu Min den Max--------")

	fmt.Println("Nhap min")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	str1 := scanner3.Text()

	fmt.Println("Nhap max")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	str2 := scanner2.Text()

    min,e1 := strconv.Atoi(str1)
	max,e2 := strconv.Atoi(str2)
	
	if e1 != nil && e2 != nil {
		fmt.Println("conversion error1:", str1)
		fmt.Println("conversion error2:", str2)
	}

	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			fmt.Println(product)
		}
	}

}
