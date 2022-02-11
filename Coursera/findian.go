package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Intialize string
	var str string
	// Scanf
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Now find indexes
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	switch {
	case !strings.HasPrefix(str, "i"):
		fmt.Printf("Not Found!")
	case !strings.Contains(str, "a"):
		fmt.Printf("Not Found!")
	case !strings.HasSuffix(str, "n"):
		fmt.Printf("Not Found!")
	default:
		fmt.Printf("Found!")
	}
}
