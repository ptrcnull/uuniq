package main

import (
	"bufio"
	"os"
)

func main() {
	history := make([]string, 0, 1024)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		found := false
		for _, existing := range history {
			if line == existing {
				found = true
			}
		}
		if !found {
			_, _ = os.Stdout.WriteString(line + "\n")
			history = append(history, line)
		}
	}
}
