package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	filter "github.com/getkin/kin-openapi/openapi3filter"
)

/*
func NewGinPetServer(petStore *PetStore, port int) *http.Server {
	swagger, err := GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic chi router
	r := gin.Default()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	r = RegisterHandlers(r, petStore)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
	return s
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	petStore := NewPetStore()

	s := NewGinPetServer(petStore, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
*/

func main() {
	var port = flag.Int("port", 4010, "Port for test HTTP server")
	var secure = flag.Bool("secure", false, "Secures the API from all authorisation problems")
	var limit = flag.Bool("limit", false, "enable limit on CreatePet and decrease the performance of CreatePetstore")
	var disablePet = flag.Bool("disablePet", false, "Disable pet endpoints")
	flag.Parse()

	// Set insecure bool
	Insecure = !*secure
	Limit = *limit
	DisablePet = *disablePet

	swagger, err := GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStore := NewPetStoreAPI()

	validatorOptions := &middleware.Options{}

	validatorOptions.Options.AuthenticationFunc = func(c context.Context, input *filter.AuthenticationInput) error {
		// Authentication is handle directly in the api.go code
		return nil
	}

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

	if Insecure {
		e.Logger.Error("API Running in insecure mode")
	} else {
		e.Logger.Error("API Running in secure mode")
	}

	// We now register our petStore above as the handler for the interface
	RegisterHandlers(e, petStore)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
