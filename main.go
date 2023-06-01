package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("in.txt")
	defer file.Close()
	readFile := bufio.NewReader(file)
	var countLines int
	for {
		_, err := readFile.ReadString('\n')
		countLines += 1
		if err != nil {
			errorStrings := fmt.Errorf("Нет строк")
			println(errorStrings.Error())
		}
		if err == io.EOF {
			errorEnd := fmt.Errorf("Конец файла")
			println(errorEnd.Error())
			break
		}

	}
	fmt.Printf("Total strings %d", countLines)
}
