package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siampudan/mytop100movie/config"
	"github.com/siampudan/mytop100movie/database"
	moviesHandler "github.com/siampudan/mytop100movie/movies/handler"
	moviesRepo "github.com/siampudan/mytop100movie/movies/repo"
	moviesUsecase "github.com/siampudan/mytop100movie/movies/usecase"
	"go.uber.org/zap"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Hello World",
	})
}

func main() {
	config.LoadConfig()
	log, _ := zap.NewDevelopment()
	client := &http.Client{}
	r := gin.Default()
	r.GET("/", Hello)
	v1 := r.Group("/v1")
	moviesRepo := moviesRepo.NewMoviesRepository(database.GetConnection(), client, log)
	moviesUsecase := moviesUsecase.NewMoviesUseCaSe(moviesRepo)

	moviesHandler.MoviesRoute(moviesUsecase, v1)
	r.Run("localhost:8080")

}
