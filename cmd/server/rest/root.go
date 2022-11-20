package rest

import (
	"fmt"
	"log"

	"github.com/getkin/kin-openapi/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
)

type IEcho interface {
	Use(middleware ...echo.MiddlewareFunc)
	Start(address string) error
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type Server struct {
	Controller controller.IController
	Echo       IEcho
	OpenAPI    *routers.Router
	Config     *configs.Config
}

func NewServer(config *configs.Config, controller *controller.Controller) *Server {
	openapi := loadOpenAPI()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(validate(e, openapi))

	server := &Server{
		Controller: controller,
		Echo:       e,
		OpenAPI:    openapi,
		Config:     config,
	}
	registerHandlers(e, server)

	return server
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%s", s.Config.Server.Port)
	if err := s.Echo.Start(addr); err != nil {
		log.Panicf("Couldn't start REST server: %s", err)
	}
}

func registerHandlers(router *echo.Echo, s *Server) {
	router.GET("/repository", s.GetAllRepositories)
	router.POST("/repository", s.CreateRepository)
	router.DELETE("/repository/:id", s.DeleteRepository)
	router.GET("/repository/:id", s.GetRepositoryByID)
	router.PUT("/repository/:id", s.UpdateRepository)
	router.POST("/repository/:id/scan", s.ScanRepository)
	router.GET("/repository/:id/scan/result", s.GetAllResultsByRepositoryID)
}
