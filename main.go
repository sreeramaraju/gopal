package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello Mr. ....?")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}
