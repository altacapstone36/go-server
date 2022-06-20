// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login and get Authorization Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Send Request Email and Password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.User"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Route Path for Get New Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Refresh Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/outpatient": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient with New Medic Record Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "GetAllPatient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date Filter Range Start (2022-02-23)",
                        "name": "date_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Date Filter Range End (2006-02-25)",
                        "name": "date_end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.OutPatientResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Add New Medic Record for Patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "New Medic Record",
                "parameters": [
                    {
                        "description": "New Medic Record",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminMedicRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/outpatient/process": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Process Medic Record by Doctor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "Process Doctor",
                "parameters": [
                    {
                        "description": "Process Medic Record by Doctor",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DoctorMedicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "GetAllPatient",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.Patient"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "CreatePatient",
                "parameters": [
                    {
                        "description": "Patient Details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch Patient Data By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "GetPatientByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.PatientDetails"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "DeletePatient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id/update": {
            "put": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "UpdatePatient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patient Details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/refresh_token": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Route Path for Get New Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass access_token Here",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "request.AdminMedicRecord": {
            "type": "object",
            "properties": {
                "complaint": {
                    "type": "string",
                    "example": "sakit perut"
                },
                "date_check": {
                    "type": "string",
                    "example": "2022-06-24"
                },
                "medical_facility_id": {
                    "type": "integer",
                    "example": 1
                },
                "medical_staff_id": {
                    "type": "integer",
                    "example": 1
                },
                "patient_id": {
                    "type": "integer",
                    "example": 1
                },
                "session_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.DoctorMedicRequest": {
            "type": "object",
            "properties": {
                "diagnose": {
                    "type": "string",
                    "example": "maag"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "prescription": {
                    "type": "string",
                    "example": "entrostop"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "alsyadahmad@holyhos.co.id"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        },
        "request.Patient": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Sumenep"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2001-04-14"
                },
                "blood_type": {
                    "type": "string",
                    "example": "A"
                },
                "code": {
                    "type": "string",
                    "example": "RM0001"
                },
                "full_name": {
                    "type": "string",
                    "example": "Faizur Ramadhan"
                },
                "gender": {
                    "type": "string",
                    "example": "Male"
                },
                "resident_registration": {
                    "type": "string",
                    "example": "8729301745162748"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MedicRecord": {
            "type": "object",
            "properties": {
                "blood_tension": {
                    "type": "integer",
                    "example": 124
                },
                "body_temp": {
                    "type": "integer",
                    "example": 34
                },
                "complaint": {
                    "type": "string",
                    "example": "Sakit Perut"
                },
                "date_check": {
                    "type": "string",
                    "example": "2022-06-17"
                },
                "diagnose": {
                    "type": "string",
                    "example": "Maag"
                },
                "doctor": {
                    "type": "string",
                    "example": "Alsyad Ahmad"
                },
                "facility": {
                    "type": "string",
                    "example": "General"
                },
                "height": {
                    "type": "integer",
                    "example": 55
                },
                "prescription": {
                    "type": "string",
                    "example": "Entrostop"
                },
                "serial_number": {
                    "type": "string",
                    "example": "RM/748/2022/0001"
                },
                "weight": {
                    "type": "integer",
                    "example": 150
                }
            }
        },
        "response.MessageData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MessageOnly": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.OutPatientResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "complaint": {
                    "type": "string"
                },
                "date_check": {
                    "type": "string"
                },
                "doctor": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "queue": {
                    "type": "integer"
                },
                "serial_number": {
                    "type": "string"
                },
                "session": {
                    "type": "string"
                }
            }
        },
        "response.Patient": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string",
                    "example": "2001-04-14"
                },
                "code": {
                    "type": "string",
                    "example": "RM0001"
                },
                "full_name": {
                    "type": "string",
                    "example": "Faizur Ramadhan"
                },
                "gender": {
                    "type": "string",
                    "example": "Male"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "resident_registration": {
                    "type": "string",
                    "example": "8729301745162748"
                }
            }
        },
        "response.PatientDetails": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Sumenep"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2001-04-14"
                },
                "blood_type": {
                    "type": "string",
                    "example": "A"
                },
                "code": {
                    "type": "string",
                    "example": "RM0001"
                },
                "full_name": {
                    "type": "string",
                    "example": "Faizur Ramadhan"
                },
                "gender": {
                    "type": "string",
                    "example": "Male"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "medic_record": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.MedicRecord"
                    }
                },
                "resident_registration": {
                    "type": "string",
                    "example": "8729301745162748"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "DR00001"
                },
                "email": {
                    "type": "string",
                    "example": "alsyadahmad@holyhos.co.id"
                },
                "facility": {
                    "type": "string",
                    "example": "General"
                },
                "full_name": {
                    "type": "string",
                    "example": "Alsyad Ahmad"
                },
                "gender": {
                    "type": "string",
                    "example": "Male"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "roles": {
                    "type": "string",
                    "example": "Doctor"
                },
                "status": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKey": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ec2-3-91-177-221.compute-1.amazonaws.com",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Holy Hospital Sever API",
	Description:      "server API for Holy Hospital Application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
