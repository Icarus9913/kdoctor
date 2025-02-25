// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "agent http server",
    "title": "http server API",
    "version": "v1"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "echo http request",
        "tags": [
          "echo"
        ],
        "summary": "echo http request",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/EchoRes"
            }
          }
        }
      }
    },
    "/healthy/liveness": {
      "get": {
        "description": "pod liveness probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/healthy/readiness": {
      "get": {
        "description": "pod readiness probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/healthy/startup": {
      "get": {
        "description": "pod startup probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/kdoctoragent": {
      "get": {
        "description": "echo http request",
        "tags": [
          "echo"
        ],
        "summary": "echo http request",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/EchoRes"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "EchoRes": {
      "description": "echo request",
      "type": "object",
      "properties": {
        "clientIp": {
          "description": "client source ip",
          "type": "string"
        },
        "otherDetail": {
          "description": "other  information",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "requestHeader": {
          "description": "request header",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "requestUrl": {
          "description": "request url",
          "type": "string"
        },
        "serverName": {
          "description": "server host name",
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "http"
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "agent http server",
    "title": "http server API",
    "version": "v1"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "echo http request",
        "tags": [
          "echo"
        ],
        "summary": "echo http request",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/EchoRes"
            }
          }
        }
      }
    },
    "/healthy/liveness": {
      "get": {
        "description": "pod liveness probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Liveness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/healthy/readiness": {
      "get": {
        "description": "pod readiness probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Readiness probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/healthy/startup": {
      "get": {
        "description": "pod startup probe for agent and controller pod",
        "tags": [
          "healthy"
        ],
        "summary": "Startup probe",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Failed"
          }
        }
      }
    },
    "/kdoctoragent": {
      "get": {
        "description": "echo http request",
        "tags": [
          "echo"
        ],
        "summary": "echo http request",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/EchoRes"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "EchoRes": {
      "description": "echo request",
      "type": "object",
      "properties": {
        "clientIp": {
          "description": "client source ip",
          "type": "string"
        },
        "otherDetail": {
          "description": "other  information",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "requestHeader": {
          "description": "request header",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "requestUrl": {
          "description": "request url",
          "type": "string"
        },
        "serverName": {
          "description": "server host name",
          "type": "string"
        }
      }
    }
  },
  "x-schemes": [
    "http"
  ]
}`))
}
