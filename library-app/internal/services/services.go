package services

import "github.com/romaxa83/mst-app/library-app/internal/repositories"

type Books interface {
}

type Services struct {
}

type Deps struct {
	Repos *repositories.Repo
}

func NewServices(deps Deps) *Services {
	return &Services{}
}
