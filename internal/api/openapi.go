package api

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// NewOpenAPIv3 instantiates the OpenAPI specification
func NewOpenAPIv3(cfg *Config) *openapi3.T {
	spec := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Password Manager",
			Description: "",
			Version:     "0.0.0",
			Contact: &openapi3.Contact{
				URL: "https://github.com/okutsen/PasswordManager",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://" + cfg.LocalAddress(),
			},
		},
	}
	spec.Components.Schemas = openapi3.Schemas{
		"Record": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("name", openapi3.NewStringSchema()).
				WithProperty("login", openapi3.NewStringSchema()).
				WithProperty("password", openapi3.NewStringSchema())),
		"Error": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("message", openapi3.NewStringSchema())),
		// User
	}
	spec.Components.RequestBodies = openapi3.RequestBodies{
		"CreateRecord": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a record.").
				WithRequired(true).
				// RecordAPI
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("id", openapi3.NewUUIDSchema()).
					WithProperty("name", openapi3.NewStringSchema()).
					WithProperty("login", openapi3.NewStringSchema()).
					WithProperty("password", openapi3.NewStringSchema())),
		},
		"UpdateRecord": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a record.").
				WithRequired(true).
				// RecordAPI
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("id", openapi3.NewUUIDSchema()).
					WithProperty("name", openapi3.NewStringSchema()).
					WithProperty("login", openapi3.NewStringSchema()).
					WithProperty("password", openapi3.NewStringSchema())),
		},
	}
	// TODO:
	// spec.Components.RequestBodies
	// spec.Components.Responses
	// spec.Paths
	return spec
}
