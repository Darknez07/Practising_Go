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
	fmt.Printf("Enter marks (0-100): ")
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
	fmt.Printf("%d", in)

	// Array
	var x [5]int = [5]int{6, 7, 2, 1, 6}
	for i, v := range x {
		fmt.Println("Dekho jaara", i, "par haii", v)
	}
	// Using slices
	// So if we don't use ... or length
	// By default this will be slice
	arr := [...]string{"abc", "def", "fck", "und", "prk", "njk"}
	s1 := arr[1:3]
	for i, v := range s1 {
		fmt.Println("Dekho ", i, " par kya mila ", v)
	}

	// Hash table
	// Key -> value pair
	// Key has to be unique
	// Hash function is assigned.
	// Maps are example of hash table

	// var mp map[string]int // intializes
	// mp = make(map[string]int) // creates empty map
	// ANother method
	mps := map[string]int{"Joe":1231, "SDjkla":12389719}

	fmt.Println(mps["Joe"])

	// Add new value
	mps["Lila"] = 123

	// Overwrite existing value
	mps["Joe"] = 124

	// Del
	delete(mps,"Joe")

	// See
	ip, val := mps["Lila"]
	fmt.Println(val) // IS present
	fmt.Println(ip) // Value

	fmt.Println(len(mps)) // Length of map

	// Loop through map
	for key,val := range mps {
		fmt.Println(key,val) // WTF
	}
}
