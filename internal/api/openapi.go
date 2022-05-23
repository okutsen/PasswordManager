package api

import(
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
	// TODO:
	// spec.Components.Schemas
	// spec.Components.RequestBodies
	// spec.Components.Responses
	// spec.Paths
	return spec
}