package repository

import "connectly/model"

func (r *repository) CreateOrUpdateUser(username, first_name, last_name string, bot_id int64, bot_name string) (model.User, error) {
	var user model.User
	q := `
    INSERT INTO users(username, first_name, last_name, bot_id, bot_name)
    VALUES($1, $2, $3, $4, $5)
      ON CONFLICT(username,bot_id) DO UPDATE SET
        first_name=excluded.first_name,
        last_name=excluded.last_name
    RETURNING id, username, first_name, last_name, bot_id, bot_name;
  `

	err := r.db.Get(&user, q, username, first_name, last_name, bot_id, bot_name)
	return user, err
}

func (r *repository) GetUser(user_id float64) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT id, username, first_name, last_name, bot_id, bot_name FROM users WHERE id=$1", user_id)
	return user, err
}

func (r *repository) GetProduct(product_id float64) (model.Product, error) {
	var product model.Product
	q := `
  SELECT id, name FROM products WHERE id=$1;
  `
	err := r.db.Get(&product, q, product_id)
	return product, err
}

func (r *repository) CreateOrUpdateReview(user_id, product_id, rate float64) error {
	q := `
    INSERT INTO reviews(user_id, product_id, rate)
    VALUES(:user_id, :product_id, :rate)
      ON CONFLICT(user_id, product_id) DO UPDATE SET
        rate=excluded.rate,
        datetime=current_timestamp;
  `

	p := map[string]interface{}{
		"user_id":    user_id,
		"product_id": product_id,
		"rate":       rate,
	}

	_, err := r.db.NamedExec(q, p)
	return err
}

func (r *repository) GetAllReviews() ([]model.Review, error) {
	var res []model.Review
	q := `
  SELECT
    r.id,
    r.rate,
    r.datetime,
    u.id as 'user.id',
    u.username as 'user.username',
    u.first_name as 'user.first_name',
    u.last_name as 'user.last_name',
    u.bot_id as 'user.bot_id',
    u.bot_name as 'user.bot_name',
    p.id as 'product.id',
    p.name as 'product.name'
  FROM reviews r
    JOIN users u on r.user_id = u.id
	  JOIN products p on r.product_id = p.id
  `
	err := r.db.Select(&res, q)
	return res, err
}
