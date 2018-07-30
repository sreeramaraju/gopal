package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello Mr. ....?")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}
