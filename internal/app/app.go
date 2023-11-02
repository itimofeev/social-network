package app

type Repository interface {
}

type Config struct {
	Repository Repository
}

type App struct {
}

func New(cfg Config) (*App, error) {
	return &App{}, nil
}
