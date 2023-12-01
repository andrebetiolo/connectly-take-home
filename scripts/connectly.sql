-- ROLLBACK;
-- BEGIN;
CREATE TABLE IF NOT EXISTS users(
	id         integer primary key autoincrement,
	username   text not null,
	first_name text not null,
	last_name  text not null,
	bot_id     text not null,
	bot_name   text not null
);

create unique index if not exists idx_users_bot_id on users (username, bot_id);

/*
INSERT INTO users(username, first_name, last_name, bot_id)
VALUES('username', 'first_name2', 'last_name2', 'bot_id')
  ON CONFLICT(username,bot_id) DO UPDATE SET
	first_name=excluded.first_name,
	last_name=excluded.last_name
RETURNING username, first_name, last_name, bot_id;

select * from users;
 */

CREATE TABLE IF NOT EXISTS users_chats(
	id         integer primary key autoincrement,
	user_id    integer not null,
	chat_id    integer not null,
	message_id integer not null,
	message    text not null,
	datetime default current_timestamp,
	FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS products(
	id			integer primary key autoincrement,
	name		text not null
);

INSERT INTO products(name) VALUES('Product A');
INSERT INTO products(name) VALUES('Product B');

CREATE TABLE IF NOT EXISTS reviews(
	id			integer primary key autoincrement,
	user_id		integer not null,
	product_id	integer not null,
	rate		integer not null,
	datetime   	datetime default current_timestamp,
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(product_id) REFERENCES products(id)
);

create index if not exists idx_reviews_user_id on reviews(user_id);
create index if not exists idx_reviews_product_id on reviews(product_id);
create unique index if not exists unq_reviews on reviews(user_id,product_id);

/*
INSERT INTO reviews(user_id, product_id, rate)
VALUES(1, 1, 3)
  ON CONFLICT(user_id, product_id) DO UPDATE SET
    rate=excluded.rate,
    datetime=current_timestamp;

SELECT
	r.id,
	u.id as 'user.id',
	u.first_name as 'user.first_name',
	u.last_name as 'user.last_name',
	u.bot_id as 'user.bot_id',
	p.id as 'product.bot_id',
	p.name as 'product.name'
FROM reviews r
	JOIN users u on r.user_id = u.id
	JOIN products p on r.product_id = p.id;

*/
--COMMIT;

-- sqlite3://path/to/database?query

go install github.com/golang-migrate/migrate@latest
