package main

import (
	"fmt"
	"strconv"

	"github.com/ishan-p/go-josephus-simulation/josephus"
)

func main() {
	var noOfSoldiers, hops int
	var noOfSoldiersStr, hopsStr string
	fmt.Printf("Enter number of soldiers: ")
	_, err := fmt.Scan(&noOfSoldiersStr)
	noOfSoldiers, err = strconv.Atoi(noOfSoldiersStr)
	if err != nil {
		fmt.Println("Please enter a valid integer")
		return
	}
	fmt.Printf("Enter number of hops: ")
	_, err = fmt.Scan(&hopsStr)
	hops, err = strconv.Atoi(hopsStr)
	if err != nil {
		fmt.Println("Please enter a valid integer")
		return
	}

	josephus.PlayJosephusSimulation(noOfSoldiers, hops)
}
