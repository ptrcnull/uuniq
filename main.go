package main

import (
	"bufio"
	"os"
)

func main() {
	history := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !history[line] {
			_, _ = os.Stdout.WriteString(line + "\n")
			history[line] = true
		}
	}
}
