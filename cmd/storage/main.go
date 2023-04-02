package main
import "fmt"

import "github.com/EvaCaufmann/storage/internal/storage"

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}