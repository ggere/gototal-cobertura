package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	totalcover "github.com/ggere/gototal-cobertura/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	percentage, err := totalcover.TotalCover(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total coverage: %.2f%%\r\n", percentage)
}
