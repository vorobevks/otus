package container

import (
	"otus/config"
)

type Container struct {
	Config *config.Config
}

func New(
	cfg *config.Config,
) Container {
	return Container{
		cfg,
	}
}
