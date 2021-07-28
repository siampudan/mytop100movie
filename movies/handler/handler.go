package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siampudan/mytop100movie/apperror"
	"github.com/siampudan/mytop100movie/movies"
)

func MoviesRoute(usecase movies.MovieUseCase, r *gin.RouterGroup) {
	handler := &Handler{
		usecase: usecase,
	}

	r.GET("/movies/:movie_id", handler.GetMovieDetail)
}

type Handler struct {
	usecase movies.MovieUseCase
}

func (h *Handler) GetMovieDetail(c *gin.Context) {
	result, err := h.usecase.GetMovieDetail(c)
	if err != nil {
		apperror.Response(c, err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
