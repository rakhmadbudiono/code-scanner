package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rakhmadbudiono/code-scanner/internal/orm"
)

func (s *Server) GetAllRepositories(ctx echo.Context) error {
	repos, err := s.Controller.GetAllRepositories()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, repos)
}

func (s *Server) CreateRepository(ctx echo.Context) error {
	repo := &orm.Repository{}
	err := ctx.Bind(repo)
	if err != nil {
		return err
	}

	repo, err = s.Controller.CreateRepository(repo)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, repo)
}

func (s *Server) DeleteRepository(ctx echo.Context) error {
	err := s.Controller.DeleteRepository(ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.NoContent(204)
}

func (s *Server) GetRepositoryByID(ctx echo.Context) error {
	repo, err := s.Controller.GetRepositoryByID(ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, repo)
}

func (s *Server) UpdateRepository(ctx echo.Context) error {
	repo := &orm.Repository{}
	repo.ID = ctx.Param("id")
	err := ctx.Bind(repo)
	if err != nil {
		return err
	}

	repo, err = s.Controller.UpdateRepository(repo)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, repo)
}

func (s *Server) ScanRepository(ctx echo.Context) error {
	err := s.Controller.ScanRepository(ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.NoContent(204)
}

func (s *Server) GetAllResultsByRepositoryID(ctx echo.Context) error {
	results, err := s.Controller.GetAllResultsByRepositoryID(ctx.Param("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, results)
}
