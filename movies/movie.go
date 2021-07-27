package movies

import (
	"github.com/gin-gonic/gin"
)

const MoviePath string = "/movie"

type Movie map[string]interface{}

type MovieRepository interface {
	GetMovieDetail(movieID int) (Movie, error)
}

type MovieUseCase interface {
	GetMovieDetail(*gin.Context) (Movie, error)
}
