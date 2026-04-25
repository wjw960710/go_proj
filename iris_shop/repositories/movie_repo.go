package repositories

import "iris_shop/datamodels"

type MovieRepository interface {
	GetMovieTitle() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieTitle() string {
	//模擬賦值給模型
	movie := &datamodels.Movie{Title: "戰爭英雄"}
	return movie.Title
}
