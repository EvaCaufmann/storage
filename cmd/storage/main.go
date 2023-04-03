package main

import (
	"fmt"
	"log"

	"github.com/EvaCaufmann/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

st.Upload("text.txt", []byte("hello"))

if err != nil {
	log.Fatal(err)
}
	fmt.Println("it uploaded", file)
}