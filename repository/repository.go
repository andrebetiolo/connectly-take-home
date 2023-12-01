package repository

import (
	"connectly/model"

	"github.com/jmoiron/sqlx"
	"github.com/yggbrazil/go-toolbox/database"
)

type Repository interface {
	CreateOrUpdateUser(username, first_name, last_name string, bot_id int64, bot_name string) (model.User, error)
	GetUser(user_id float64) (model.User, error)
	GetProduct(product_id float64) (model.Product, error)
	CreateOrUpdateReview(user_id, product_id, rate float64) error
	GetAllReviews() ([]model.Review, error)
}

type repository struct {
	db     *sqlx.DB
	config Config
}

type Config struct {
	Type         string
	PathToDBFile string
}

func New(c Config) (Repository, error) {
	cfg := database.Config{
		Type:         c.Type,
		PathToDBFile: c.PathToDBFile,
	}

	db, err := database.ConnectByConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:     db,
		config: c,
	}, nil
}
