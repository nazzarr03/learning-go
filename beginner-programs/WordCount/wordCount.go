package word_count

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	text := strings.Join(os.Args[1:], " ")
	array := strings.Fields(text)

	count := make(map[string]int)

	fmt.Println(array)

	for _, word := range array {
		count[word]++
	}

	for word, c := range count {
		fmt.Println(word, c)
	}
}
