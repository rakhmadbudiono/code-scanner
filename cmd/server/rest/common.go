package rest

import (
	"errors"
	"fmt"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/labstack/echo/v4"
)

func loadOpenAPI() *routers.Router {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./api/openapi.yaml")
	if err != nil {
		log.Fatalf("Couldn't load openapi spec: %s", err)
	}

	router, err := legacy.NewRouter(doc)
	if err != nil {
		log.Fatalf("Couldn't create openapi router: %s", err)
	}

	return &router
}

func validateRequest(ctx echo.Context, r routers.Router) error {
	req := ctx.Request()
	route, params, err := r.FindRoute(req)
	if err != nil {
		message := fmt.Sprintf("route not found: %s", err.Error())
		return errors.New(message)
	}

	validation := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: params,
		Route:      route,
	}

	return openapi3filter.ValidateRequest(ctx.Request().Context(), validation)
}

func validate(e *echo.Echo, r *routers.Router) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := validateRequest(c, *r)
			if err != nil {
				return echo.ErrBadRequest
			}

			return next(c)
		}
	}
}
