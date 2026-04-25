package controllers

import (
	"iris_shop/repositories"
	"iris_shop/services"

	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepo)
	movieResult := movieService.ShowMovieTitle()
	return mvc.View{
		Name: "movie.html",
		Data: movieResult,
	}
}
