package main

import (
	"bufio"
	"os"
	"strings"

	srv "github.com/anujkumarthakur/service1"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	words := strings.Fields(inputLine)
	srv.ServiceFirst(words)

}
