package converter

import (
    "encoding/json"
    "io/ioutil"

    "github.com/getkin/kin-openapi/openapi3"
)

type OpenAPI31 struct {
    Openapi    string                 `json:"openapi"`
    Info       Info                   `json:"info"`
    Components Components31           `json:"components"`
    Paths      map[string]interface{} `json:"paths"`
}

type OpenAPI30 struct {
    Openapi    string                 `json:"openapi"`
    Info       Info                   `json:"info"`
    Components Components30           `json:"components"`
    Paths      map[string]interface{} `json:"paths"`
}

type Info struct {
    Title   string `json:"title"`
    Version string `json:"version"`
}

type Components31 struct {
    SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
    Schemas         map[string]Schema31       `json:"schemas"`
}

type Components30 struct {
    SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
    Schemas         map[string]Schema30       `json:"schemas"`
}

type SecurityScheme struct {
    Type string `json:"type"`
    Name string `json:"name"`
    In   string `json:"in"`
}

type Schema31 struct {
    Type                 *string               `json:"type,omitempty"`
    Properties           map[string]Property31 `json:"properties,omitempty"`
    Items                *Schema31             `json:"items,omitempty"`
    AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
    Enum                 []interface{}         `json:"enum,omitempty"`
    Required             []string              `json:"required,omitempty"`
    OneOf                []*Schema31           `json:"oneOf,omitempty"`
    AnyOf                []*Schema31           `json:"anyOf,omitempty"`
    AllOf                []*Schema31           `json:"allOf,omitempty"`
    Format               string                `json:"format,omitempty"`
    Title                string                `json:"title,omitempty"`
}

type Schema30 struct {
    Type                 string                `json:"type,omitempty"`
    Properties           map[string]Property30 `json:"properties,omitempty"`
    Items                *Schema30             `json:"items,omitempty"`
    AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
    Enum                 []string              `json:"enum,omitempty"`
    Required             []string              `json:"required,omitempty"`
    AnyOf                []*Schema30           `json:"anyOf,omitempty"`
    Format               string                `json:"format,omitempty"`
    Title                string                `json:"title,omitempty"`
}

type Property31 struct {
    Type                 *string               `json:"type,omitempty"`
    Format               string                `json:"format,omitempty"`
    Enum                 []interface{}         `json:"enum,omitempty"`
    Properties           map[string]Property31 `json:"properties,omitempty"`
    Items                *Property31           `json:"items,omitempty"`
    AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
    OneOf                []*Property31         `json:"oneOf,omitempty"`
    AnyOf                []*Property31         `json:"anyOf,omitempty"`
    AllOf                []*Property31         `json:"allOf,omitempty"`
}

type Property30 struct {
    Type                 string                `json:"type,omitempty"`
    Format               string                `json:"format,omitempty"`
    Enum                 []string              `json:"enum,omitempty"`
    Properties           map[string]Property30 `json:"properties,omitempty"`
    Items                *Property30           `json:"items,omitempty"`
    AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
    AnyOf                []*Property30         `json:"anyOf,omitempty"`
}

func Convert(inputFile, outputFile string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return err
    }

    var openAPI31 OpenAPI31
    err = json.Unmarshal(data, &openAPI31)
    if err != nil {
        return err
    }

    loader := openapi3.NewLoader()
    spec, err := loader.LoadFromData(data)
    if err != nil {
        return err
    }

    openAPI30 := convertToOpenAPI30(spec)

    outputData, err := json.MarshalIndent(openAPI30, "", "  ")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(outputFile, outputData, 0644)
    if err != nil {
        return err
    }

    return nil
}

func convertToOpenAPI30(spec *openapi3.T) OpenAPI30 {
    openAPI30 := OpenAPI30{
        Openapi: "3.0.3",
        Info: Info{
            Title:   spec.Info.Title,
            Version: spec.Info.Version,
        },
        Components: Components30{
            SecuritySchemes: make(map[string]SecurityScheme),
            Schemas:         make(map[string]Schema30),
        },
        Paths: convertPaths(spec.Paths),
    }

    for key, scheme := range spec.Components.SecuritySchemes {
        openAPI30.Components.SecuritySchemes[key] = SecurityScheme{
            Type: scheme.Value.Type,
            Name: scheme.Value.Name,
            In:   scheme.Value.In,
        }
    }

    for key, schema := range spec.Components.Schemas {
        openAPI30.Components.Schemas[key] = *convertSchema31To30(schema.Value)
    }

    return openAPI30
}

func convertPaths(paths openapi3.Paths) map[string]interface{} {
    result := make(map[string]interface{})
    for k, v := range paths {
        result[k] = v
    }
    return result
}

func convertSchema31To30(schema31 *openapi3.Schema) *Schema30 {
    schema30 := &Schema30{
        Type:                 derefString(schema31.Type),
        Properties:           make(map[string]Property30),
        AdditionalProperties: schema31.AdditionalProperties.Has,
        Enum:                 convertEnum31To30(schema31.Enum),
        Required:             schema31.Required,
        AnyOf:                []*Schema30{},
        Format:               schema31.Format,
        Title:                schema31.Title,
    }

    for key, property31 := range schema31.Properties {
        schema30.Properties[key] = *convertProperty31To30(property31.Value)
    }

    if schema31.Items != nil {
        schema30.Items = convertSchema31To30(schema31.Items.Value)
    }

    for _, anyOfSchema := range schema31.AnyOf {
        schema30.AnyOf = append(schema30.AnyOf, convertSchema31To30(anyOfSchema.Value))
    }

    return schema30
}

func convertEnum31To30(enum31 []interface{}) []string {
    var enum30 []string
    for _, e := range enum31 {
        if str, ok := e.(string); ok {
            enum30 = append(enum30, str)
        }
    }
    return enum30
}

func convertProperty31To30(property31 *openapi3.Schema) *Property30 {
    property30 := &Property30{
        Type:                 derefString(property31.Type),
        Format:               property31.Format,
        Enum:                 convertEnum31To30(property31.Enum),
        Properties:           make(map[string]Property30),
        AdditionalProperties: property31.AdditionalProperties.Has,
        AnyOf:                []*Property30{},
    }

    for key, subProperty31 := range property31.Properties {
        property30.Properties[key] = *convertProperty31To30(subProperty31.Value)
    }

    if property31.Items != nil {
        property30.Items = convertProperty31To30(property31.Items.Value)
    }

    for _, anyOfProperty := range property31.AnyOf {
        property30.AnyOf = append(property30.AnyOf, convertProperty31To30(anyOfProperty.Value))
    }

    return property30
}

func derefString(str *string) string {
    if str == nil {
        return ""
    }
    return *str
}
