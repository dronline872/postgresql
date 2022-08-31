/* 

Интернет-магазин бытовой техники

use-case:
просмотр товара -> 
добавление в корзину -> 
изменение количество товаров/ удаление из корзины ->
добавление в заказы(оформление) ->
просмотр заказа ->
изменение статуса заказа
*/

-- Пользователи
CREATE TABLE users( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name Character Varying,
    email Character Varying
);

-- Товары
CREATE TABLE products( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name Character Varying,
	price Real DEFAULT 0 NOT NULL,
	sort Integer DEFAULT 0 NOT NULL,
	active Boolean DEFAULT 'true' NOT NULL,
	description Text,
	img Character Varying
);

-- Корзины пользователей
CREATE TABLE carts ( 
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id Integer NOT NULL,
	product_id Integer NOT NULL,
	product_count Integer DEFAULT 0 NOT NULL,
	created_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	updated_at Timestamp Without Time Zone DEFAULT now() NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (product_id) REFERENCES products (id)
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
	FOREIGN KEY (product_id) REFERENCES products (id)
);