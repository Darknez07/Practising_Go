package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"bufio"
	"strings"
)
// Create a struct for name
type Name struct{
	fname string
	lname string
}

func main(){
	// Enter for filename
	var fn string
	fmt.Printf("Enter the filename (.txt): ")
	fmt.Scan(&fn)
	//  Read the file
	data,e := ioutil.ReadFile(fn)
	if e != nil{
		// Log out the error
		log.Fatal("Error occured")
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	// Make slice of struct Name
	arr := make([]Name,0)
	// The scanner is sacning file line by line
	for scanner.Scan() {
		vals := strings.Split(scanner.Text()," ")
		// Storing names as Name struct
		arr = append(arr,Name{fname:vals[0],lname:vals[1]})
	}
	// Looping over the slice of struct Name
	for _,v := range arr{
		fmt.Printf("%s ",v.fname)
		fmt.Printf("%s\n",v.lname)
	}
}