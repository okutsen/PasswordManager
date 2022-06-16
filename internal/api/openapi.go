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
	// TODO: use openapi3gen to generate schemas from go structs
	spec.Components.Schemas = openapi3.Schemas{
		// TODO: add required properties
		"Record": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("name", openapi3.NewStringSchema()).
				WithProperty("login", openapi3.NewStringSchema()).
				WithProperty("password", openapi3.NewStringSchema())),
		// TODO: add User schema
		"Error": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("message", openapi3.NewStringSchema())),
	}
	spec.Components.Parameters = openapi3.ParametersMap{
		"RecordIDPathParam": &openapi3.ParameterRef{
			Value: openapi3.NewPathParameter(IDPPN).
				WithSchema(openapi3.NewUUIDSchema()),
		},
		"CorrelationIDHeaderParam": &openapi3.ParameterRef{
			Value: openapi3.NewHeaderParameter(CorrelationIDHPN).
				WithDescription("Correlation id").
				WithSchema(openapi3.NewUUIDSchema()),
		},
	}
	spec.Components.RequestBodies = openapi3.RequestBodies{
		"CreateRecordRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a record.").
				WithRequired(true).
				WithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/Record",
				}),
		},
		"UpdateRecordRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a record.").
				WithRequired(true).
				WithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/Record",
				}),
		},
	}
	spec.Components.Responses = openapi3.Responses{
		"ListRecordsResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returns back all records.").
				WithJSONSchema(&openapi3.Schema{
					Type:  openapi3.TypeArray,
					Items: openapi3.NewSchemaRef("#/components/schemas/Record", nil),
				}),
		},
		"RecordResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returns back successfully found or created record.").
				WithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/Record",
				}),
		},
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/Error",
				}),
		},
	}
	spec.Paths = openapi3.Paths{
		"/records": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "ListRecords",
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ListRecordsResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Post: &openapi3.Operation{
				OperationID: "CreateRecord",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateRecordRequest",
				},
				Responses: openapi3.Responses{
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/RecordResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateRecord",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateRecordRequest",
				},
				Responses: openapi3.Responses{
					"202": &openapi3.ResponseRef{
						Ref: "#/components/responses/RecordResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
		"/records/{" + IDPPN + "}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "GetRecord",
				Parameters: []*openapi3.ParameterRef{{
					Ref: "#/components/parameters/RecordIDPathParam",
				}},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/RecordResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteRecord",
				Parameters: []*openapi3.ParameterRef{{
					Ref: "#/components/parameters/RecordIDPathParam",
				}},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Record deleted"),
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
	}
	return spec
}
