package repo

import (
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/siampudan/mytop100movie/helper"
	"github.com/siampudan/mytop100movie/movies"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewMoviesRepository(db *pg.DB, client *http.Client, log *zap.Logger) *MoviesRepository {
	return &MoviesRepository{
		db:     db,
		client: client,
		log:    log,
	}
}

type MoviesRepository struct {
	db     *pg.DB
	client *http.Client
	log    *zap.Logger
}

func (repo *MoviesRepository) GetMovieDetail(movieID int) (movies.Movie, error) {
	URL := viper.GetString("tmdb.endPoint") + movies.MoviePath + "/" + strconv.Itoa(movieID) + "?api_key=" + viper.GetString("tmdb.apiKey")
	res, err := repo.client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	result, err := helper.ParseResponseBodyToMap(res)
	if err != nil {
		return nil, err
	}

	return result, nil
}
