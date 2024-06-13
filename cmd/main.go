package main

import (
    "fmt"
    "os"

    "github.com/jbcom/openapi-31-to-30-converter/pkg/converter"
    "github.com/jbcom/openapi-31-to-30-converter/pkg/validator"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <input-file> <output-file>")
        return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    err := converter.Convert(inputFile, outputFile)
    if err != nil {
        fmt.Println("Error during conversion:", err)
        return
    }

    fmt.Println("Generation completed successfully.")
    validator.ValidateOpenAPI(outputFile)
}
