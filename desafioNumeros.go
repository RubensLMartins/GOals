package main

import "fmt"

func main() {
	var lst1, lst2 []int
	for i := 1; i <= 100; i++ {
		if i%3 == 0 || i%5 == 0 {
			switch {
			case i%3 == 0:
				lst1 = append(lst1[:], i)
				fmt.Println("Pin")

			case i%5 == 0:
				lst2 = append(lst2[:], i)
				fmt.Println("Pan")
			}
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("A lista múltiplos de 3: ", lst1)
	fmt.Println("A lista múltiplos de 5: ", lst2)
}
