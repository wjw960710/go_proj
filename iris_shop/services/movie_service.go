package services

import (
	"fmt"
	"iris_shop/repositories"
)

type MovieService interface {
	ShowMovieTitle() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieTitle() string {
	title := "我們取到的電影名稱為 " + m.repo.GetMovieTitle()
	fmt.Println(title)
	return title
}
