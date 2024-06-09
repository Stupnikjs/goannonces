package main

// audio
type Track struct {
	ID             int32
	Name           string
	SelectionCount int
	PlayCount      int
	StoreURL       string
	Tags           []string
}
