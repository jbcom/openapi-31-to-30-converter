package validator

import (
    "fmt"
    "io/ioutil"

    "github.com/xeipuuv/gojsonschema"
)

func ValidateOpenAPI(filePath string) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file for validation:", err)
        return
    }

    loader := gojsonschema.NewStringLoader(string(data))
    schemaLoader := gojsonschema.NewReferenceLoader("https://spec.openapis.org/oas/3.0/schema/2021-09-28")

    result, err := gojsonschema.Validate(schemaLoader, loader)
    if err != nil {
        fmt.Println("Error validating OpenAPI document:", err)
        return
    }

    if result.Valid() {
        fmt.Println("The OpenAPI document is valid.")
    } else {
        fmt.Printf("The OpenAPI document is not valid. See errors:\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}
