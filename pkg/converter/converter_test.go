package converter

import (
    "io/ioutil"
    "testing"
)

func TestConvert(t *testing.T) {
    inputFile := "testdata/openapi31.json"
    outputFile := "testdata/openapi30.json"

    err := Convert(inputFile, outputFile)
    if err != nil {
        t.Fatalf("Conversion failed: %v", err)
    }

    outputData, err := ioutil.ReadFile(outputFile)
    if err != nil {
        t.Fatalf("Error reading output file: %v", err)
    }

    if len(outputData) == 0 {
        t.Fatalf("Output file is empty")
    }
}
