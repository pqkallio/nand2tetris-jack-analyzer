package main

import (
	"log"
	"os"

	"github.com/pqkallio/nand2tetris-jack-compiler/tokenizer"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("please provide only the file or folder to compile")
	}

	fn := os.Args[1]

	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("error opening file %s: %s", fn, err.Error())
	}

	t := tokenizer.New(f)

	t.Advance()

	for t2 := t.Token(); t2.Type != tokenizer.EOF; t2 = t.Token() {
		log.Printf("%+v", t.Token())
		t.Advance()
	}
}
