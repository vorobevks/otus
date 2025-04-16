package container

import (
	"otus/config"
	"otus/internal/repository"
)

type Container struct {
	Config     *config.Config
	Repository *repository.Repository
}

func New(
	cfg *config.Config,
	repository *repository.Repository,
) Container {
	return Container{
		cfg,
		repository,
	}
}
