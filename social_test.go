package main

import (
	"log"
	"testing"
)

func TestTestMe(t *testing.T) {
	log.Println("Testing TestMe function...")
	if i := TestMe(); i != 1 {
		t.Fatal("Failed test. Expected 1, got", i)
	}
}
