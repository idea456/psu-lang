package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(reader *bufio.Reader) string {
	t, _ := reader.ReadString('\n')
	return strings.TrimSpace(t)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	itpr := NewInterpreter()
	fmt.Print("PSU Language | psuc 1.0.0\n")
	fmt.Print("Type exit to exit the program or press Ctrl-D.\n")

	var s *Scanner
	fmt.Print(">>> ")
	line := getInput(reader)
	for ; !strings.EqualFold("exit", line); line = getInput(reader) {
		if strings.Split(line, " ")[0] == "file" {
			bytes, err := os.ReadFile(strings.Split(line, " ")[1])
			if err != nil {
				fmt.Println("Error, file not found!")
				fmt.Print(">>> ")
				continue
			}
			s = NewScanner(string(bytes))
		} else {
			s = NewScanner(line)
		}
		arr := s.Scan()
		// fmt.Println(arr)
		parser := NewParser(arr)
		stmts := parser.Parse()

		// fmt.Printf("Type: %#v\n", stmts[0])
		itpr.Interpret(stmts)
		fmt.Print(">>> ")
	}
	fmt.Print("Bai bai!\n")
}
