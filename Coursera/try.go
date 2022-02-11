package main

import (
	"fmt"
	"log"
)

func main() {
	type Grades int
	const (
		F Grades = iota
		E
		D
		C
		B
		A
	)
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err)
	}
	if num > 90 {
		fmt.Printf("%d\n", A)
	} else if num <= 90 && num > 80 {
		fmt.Printf("%d\n", B)
	} else if num <= 80 && num > 70 {
		fmt.Printf("%d\n", C)
	} else if num <= 70 && num > 60 {
		fmt.Printf("%d\n", D)
	} else if num <= 60 && num > 50 {
		fmt.Printf("%d\n", E)
	} else {
		fmt.Printf("%d\n", F)
	}
	switch {
	case num > 90:
		fmt.Printf("%d\n", A)
	case num <= 90 && num > 80:
		fmt.Printf("%d\n", B)
	}
	in := 3 % (1 + 5)
	fmt.Printf("%d",in)

	// Array
	var x [5]int = [5]int {6,7,2,1,6}
	for i, v :=range x{
		fmt.Println("Dekho jaara",i, "par haii",v)
	}
}
