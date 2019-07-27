// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/adasiunas/bello-auth/apimodel"
	"github.com/adasiunas/bello-auth/business"
	middlew "github.com/adasiunas/bello-auth/restapi/middleware"
	"github.com/adasiunas/bello-auth/restapi/operations/user"
	v1 "github.com/adasiunas/bello-auth/restapi/v1"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"os"

	"github.com/adasiunas/bello-auth/restapi/operations"
	"github.com/adasiunas/bello-auth/restapi/operations/status"
)

//go:generate swagger generate server --target ../../bello-auth --name BelloAppAuth --spec ../swagger.json --model-package apimodel

func configureFlags(api *operations.BelloAppAuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BelloAppAuthAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	f := business.NewFactory()

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = middlew.BearerAuthentication

	api.StatusServiceStatusHandler = status.ServiceStatusHandlerFunc(func(params status.ServiceStatusParams) middleware.Responder {
		hash := os.Getenv("GIT_HASH")
		return status.NewServiceStatusOK().WithPayload(&apimodel.StatusResponse{hash})
	})

	api.UserRegisterUserV1Handler = user.RegisterUserV1HandlerFunc(v1.RegisterUser(f))

	api.GetResourceHandler = operations.GetResourceHandlerFunc(func(params operations.GetResourceParams, principal interface{}) middleware.Responder {
		fmt.Println(principal)
		return operations.NewGetResourceOK().WithPayload(&operations.GetResourceOKBody{Text:"Works like a charm"})
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
