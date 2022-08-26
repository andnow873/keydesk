// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "swagger": "2.0",
  "info": {
    "title": "Keykeeper server",
    "version": "1.0.0"
  },
  "basePath": "/user",
  "paths": {
    "/list": {
      "get": {
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "A list of users.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "UserID",
        "UserName"
      ],
      "properties": {
        "LastVisitASCountry": {
          "type": "string",
          "format": "date-time"
        },
        "LastVisitASName": {
          "type": "string",
          "format": "date-time"
        },
        "LastVisitSubnet": {
          "type": "string",
          "format": "date-time"
        },
        "Problems": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ThrottlingTill": {
          "type": "string",
          "format": "date-time"
        },
        "UserID": {
          "type": "string"
        },
        "UserName": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Keykeeper server",
    "version": "1.0.0"
  },
  "basePath": "/user",
  "paths": {
    "/list": {
      "get": {
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "A list of users.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "UserID",
        "UserName"
      ],
      "properties": {
        "LastVisitASCountry": {
          "type": "string",
          "format": "date-time"
        },
        "LastVisitASName": {
          "type": "string",
          "format": "date-time"
        },
        "LastVisitSubnet": {
          "type": "string",
          "format": "date-time"
        },
        "Problems": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ThrottlingTill": {
          "type": "string",
          "format": "date-time"
        },
        "UserID": {
          "type": "string"
        },
        "UserName": {
          "type": "string"
        }
      }
    }
  }
}`))
}
