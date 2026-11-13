package main

import "fmt"

func main() {
	s := NewStore()

	words := []string{"go", "is", "fast", "go", "is", "go", "simple", "fast"}
	for _, w := range words {
		s.Add(w)
	}

	fmt.Println("Count of 'go':", s.Count("go"))
	fmt.Println("Count of 'is':", s.Count("is"))
	fmt.Println("Count of 'java':", s.Count("java"))

	fmt.Println("Top 3 words:", s.TopN(3))
}
