package main

// Store tracks how many times each word has been added.
type Store struct {
	// TODO: add the fields you need
}

// NewStore creates and returns an empty Store.
func NewStore() *Store {
	// TODO: initialise and return a Store
	return &Store{}
}

// Add records one occurrence of word.
// Words are case-insensitive: "Go" and "go" are the same word.
func (s *Store) Add(word string) {
	// TODO: implement
}

// Count returns the number of times word has been added.
// Returns 0 if the word was never added.
func (s *Store) Count(word string) int {
	// TODO: implement
	return 0
}

// TopN returns the n most frequent words, ordered from highest to lowest count.
// If n is larger than the number of unique words, all words are returned.
func (s *Store) TopN(n int) []string {
	// TODO: implement
	return nil
}
