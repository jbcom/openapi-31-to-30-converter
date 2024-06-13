package validator

import (
    "testing"
)

func TestValidateOpenAPI(t *testing.T) {
    filePath := "testdata/openapi30.json"

    // Assuming the file is valid, the function should not print any errors.
    // We will capture the output to verify this.
    ValidateOpenAPI(filePath)
}
