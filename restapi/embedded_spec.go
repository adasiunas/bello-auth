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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API for bello app authentication",
    "title": "Bello app auth",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/status": {
      "get": {
        "description": "Used for checking service status and code hash",
        "tags": [
          "Status"
        ],
        "operationId": "serviceStatus",
        "responses": {
          "200": {
            "description": "Status response",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/resource": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Test auth resource",
        "operationId": "getResource",
        "responses": {
          "200": {
            "schema": {
              "type": "object",
              "properties": {
                "text": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        }
      }
    },
    "/v1/user": {
      "post": {
        "description": "User registration endpoint",
        "tags": [
          "User"
        ],
        "operationId": "registerUserV1",
        "parameters": [
          {
            "name": "Register",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RegistrationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully registered",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/user/login": {
      "post": {
        "description": "Login user",
        "tags": [
          "User"
        ],
        "operationId": "loginUser",
        "parameters": [
          {
            "name": "Login",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully registered",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/user/refresh": {
      "post": {
        "description": "Refresh access token",
        "tags": [
          "User"
        ],
        "operationId": "refreshAccessToken",
        "parameters": [
          {
            "name": "Token",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RefreshRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User's token was successfully refreshed",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "enum": [
            "InternalServerError",
            "Unauthorized",
            "NotFound",
            "UserAlreadyExists",
            "IncorrectCredentials",
            "BadRequest"
          ]
        },
        "type": {
          "type": "string"
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "x-nullable": false
        },
        "password": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "RefreshRequest": {
      "type": "object",
      "required": [
        "refreshToken"
      ],
      "properties": {
        "refreshToken": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "RegistrationRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "x-nullable": false
        },
        "password": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "StatusResponse": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string"
        }
      }
    },
    "TokenResponse": {
      "properties": {
        "accessToken": {
          "type": "string",
          "x-nullable": false
        },
        "refreshToken": {
          "type": "string",
          "x-nullable": false
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Status endpoints",
      "name": "Status"
    },
    {
      "description": "User related operation endpoints",
      "name": "User"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API for bello app authentication",
    "title": "Bello app auth",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/status": {
      "get": {
        "description": "Used for checking service status and code hash",
        "tags": [
          "Status"
        ],
        "operationId": "serviceStatus",
        "responses": {
          "200": {
            "description": "Status response",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/resource": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Test auth resource",
        "operationId": "getResource",
        "responses": {
          "200": {
            "schema": {
              "type": "object",
              "properties": {
                "text": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        }
      }
    },
    "/v1/user": {
      "post": {
        "description": "User registration endpoint",
        "tags": [
          "User"
        ],
        "operationId": "registerUserV1",
        "parameters": [
          {
            "name": "Register",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RegistrationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully registered",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/user/login": {
      "post": {
        "description": "Login user",
        "tags": [
          "User"
        ],
        "operationId": "loginUser",
        "parameters": [
          {
            "name": "Login",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully registered",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/v1/user/refresh": {
      "post": {
        "description": "Refresh access token",
        "tags": [
          "User"
        ],
        "operationId": "refreshAccessToken",
        "parameters": [
          {
            "name": "Token",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RefreshRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User's token was successfully refreshed",
            "schema": {
              "$ref": "#/definitions/TokenResponse"
            }
          },
          "400": {
            "description": "Bad request payload",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Service down",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "enum": [
            "InternalServerError",
            "Unauthorized",
            "NotFound",
            "UserAlreadyExists",
            "IncorrectCredentials",
            "BadRequest"
          ]
        },
        "type": {
          "type": "string"
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "x-nullable": false
        },
        "password": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "RefreshRequest": {
      "type": "object",
      "required": [
        "refreshToken"
      ],
      "properties": {
        "refreshToken": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "RegistrationRequest": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "x-nullable": false
        },
        "password": {
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "StatusResponse": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string"
        }
      }
    },
    "TokenResponse": {
      "properties": {
        "accessToken": {
          "type": "string",
          "x-nullable": false
        },
        "refreshToken": {
          "type": "string",
          "x-nullable": false
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Status endpoints",
      "name": "Status"
    },
    {
      "description": "User related operation endpoints",
      "name": "User"
    }
  ]
}`))
}
