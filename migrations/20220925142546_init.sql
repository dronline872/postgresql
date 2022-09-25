-- +goose Up
-- +goose StatementBegin
-- Пользователи
CREATE TABLE users( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name Character Varying,
    email Character Varying,
	CONSTRAINT email_validate CHECK(email ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$')
);
-- Товары
CREATE TABLE products( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name Character Varying,
	price Real NOT NULL,
	sort Integer DEFAULT 0 NOT NULL,
	active Boolean DEFAULT 'true' NOT NULL,
	description Text,
	img Character Varying,
	CONSTRAINT not_zero_price CHECK (price > 0)
);
-- Корзины пользователей
CREATE TABLE carts ( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id Integer NOT NULL,
	product_id Integer NOT NULL,
	product_count Integer NOT NULL,
	created_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	updated_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (product_id) REFERENCES products (id),
	CONSTRAINT not_zero_product_count CHECK (product_count > 0)
);
-- Заказы пользователей
CREATE TABLE orders ( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	phone Character Varying,
	fio Character Varying,
	email Character Varying,
	order_comment Text,
	order_status Character Varying,
    user_id Integer NOT NULL,
	created_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	updated_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id)
);
-- Продукты в заказах пользователей
CREATE TABLE orders_products ( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	product_id Integer NOT NULL,
	product_count Integer NOT NULL,
	order_id Integer NOT NULL,
    created_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	updated_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	FOREIGN KEY (order_id) REFERENCES orders (id),
	FOREIGN KEY (product_id) REFERENCES products (id),
	CONSTRAINT not_zero_product_count CHECK (product_count > 0)
);
-- Индексы
-- Индекс для products
CREATE INDEX products_name_idx ON products USING btree (name text_pattern_ops);
-- Индекс для users
CREATE INDEX users_name_email_idx ON users (email, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- удаление индексов
DROP INDEX users_name_email_idx, products_name_idx;
-- удаление таблиц
DROP TABLE orders_products, orders, carts, products, users;
-- +goose StatementEnd
