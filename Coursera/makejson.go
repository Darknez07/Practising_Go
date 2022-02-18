package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Enter the name
	fmt.Printf("Enter the name: ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("Error Found")
	}
	// Enter the address
	fmt.Printf("Enter the address: ")
	add, er := in.ReadString('\n')
	if er != nil {
		log.Fatal("Error Found")
	}
	// Create a map
	mps := map[string]string{"Name": name, "Address": add}
	// Marshall the map to json object.
	js, e := json.Marshal(mps)
	if e != nil {
		log.Fatal("Error Found")
	}
	// Print it in json format.
	fmt.Printf(string(js))
}
