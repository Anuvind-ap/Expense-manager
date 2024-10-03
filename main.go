package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	dir, err := os.Open("statements")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		pdfPath := "statements/" + filename
		err := api.ExtractContentFile(pdfPath, "output/extracted_text.txt", nil, nil)
		if err != nil {
			log.Fatalf("Error extracting text from PDF: %s\n", err)
		}

		fmt.Println("Text extraction complete! Check output/extracted_text.txt for results.")
	}
}
