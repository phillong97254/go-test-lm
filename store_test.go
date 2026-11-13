package main

import (
	"testing"
)

func TestCountZeroForUnseenWord(t *testing.T) {
	s := NewStore()
	if got := s.Count("go"); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestCountAfterAdd(t *testing.T) {
	s := NewStore()
	s.Add("go")
	s.Add("go")
	s.Add("go")
	if got := s.Count("go"); got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
}

func TestCaseInsensitive(t *testing.T) {
	s := NewStore()
	s.Add("Go")
	s.Add("GO")
	s.Add("go")
	if got := s.Count("go"); got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
}

func TestTopNOrder(t *testing.T) {
	s := NewStore()
	words := []string{"go", "go", "go", "rust", "rust", "python"}
	for _, w := range words {
		s.Add(w)
	}

	top := s.TopN(2)
	if len(top) != 2 {
		t.Fatalf("expected 2 results, got %d", len(top))
	}
	if top[0] != "go" {
		t.Errorf("expected 'go' to be first, got %q", top[0])
	}
	if top[1] != "rust" {
		t.Errorf("expected 'rust' to be second, got %q", top[1])
	}
}

func TestTopNMoreThanAvailable(t *testing.T) {
	s := NewStore()
	s.Add("go")
	s.Add("rust")

	top := s.TopN(10)
	if len(top) != 2 {
		t.Errorf("expected 2 results, got %d", len(top))
	}
}

func TestTopNEmpty(t *testing.T) {
	s := NewStore()
	top := s.TopN(3)
	if len(top) != 0 {
		t.Errorf("expected empty result, got %v", top)
	}
}

func TestMultipleWords(t *testing.T) {
	s := NewStore()
	s.Add("fast")
	s.Add("simple")
	s.Add("fast")

	if s.Count("fast") != 2 {
		t.Errorf("expected fast=2, got %d", s.Count("fast"))
	}
	if s.Count("simple") != 1 {
		t.Errorf("expected simple=1, got %d", s.Count("simple"))
	}
}
