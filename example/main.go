package main

import (
	"fmt"

	mimedb "github.com/sniperkit/mimedb/pkg"
)

func main() {
	fmt.Println("mime-db - simple example...")

	extension := "yaml"

	if entry, ok := mimedb.DB[extension]; ok {
		fmt.Println("entry.ContentType=", entry.ContentType, "entry.Compressible=", entry.Compressible)
	}
}
