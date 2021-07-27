package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siampudan/mytop100movie/apperror"
	"github.com/siampudan/mytop100movie/constant"
	"github.com/siampudan/mytop100movie/movies"
)

func NewMoviesUseCaSe(repo movies.MovieRepository) *MovieUseCase {
	return &MovieUseCase{
		repo: repo,
	}
}

type MovieUseCase struct {
	repo movies.MovieRepository
}

func (uc *MovieUseCase) GetMovieDetail(c *gin.Context) (movies.Movie, error) {
	movieID, _ := strconv.Atoi(c.Param("movie_id"))

	result, err := uc.repo.GetMovieDetail(movieID)
	if err != nil {
		return nil, apperror.New(http.StatusBadRequest, constant.ErrorRequest)
	}
	return result, nil
}
