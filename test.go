package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	URL := "https://localhost:8080"
	fmt.Println(URL[:rand.IntN(8)] + strconv.Itoa(rand.IntN(256)))
}
