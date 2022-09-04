-- Создание пользователей
INSERT INTO users (
    name,
    email
)
VALUES
    ('Пользователь 1', 'test1@test.com'),
    ('Пользователь 2', 'test2@test.com'),
    ('Пользователь 3', 'test3@test.com'),
    ('Пользователь 4', 'test4@test.com'),
    ('Пользователь 5', 'test5@test.com'),
    ('Пользователь 6', 'test6@test.com'),
    ('Пользователь 7', 'test7@test.com');

-- Создание товаров
INSERT INTO products (
    name,
	price,
	sort,
	description,
	img
)
VALUES
    ('Товар 1', 100.0, 100, 'Большое описание товара', './images/img1.jpg'),
    ('Товар 2', 200.0, 100, 'Большое описание товара', './images/img2.jpg'),
    ('Товар 3', 300.0, 100, 'Большое описание товара', './images/img3.jpg'),
    ('Товар 4', 400.0, 100, 'Большое описание товара', './images/img4.jpg'),
    ('Товар 5', 500.0, 100, 'Большое описание товара', './images/img5.jpg'),
    ('Товар 6', 600.0, 100, 'Большое описание товара', './images/img6.jpg'),
    ('Товар 7', 700.0, 100, 'Большое описание товара', './images/img7.jpg'),
    ('Товар 8', 800.0, 100, 'Большое описание товара', './images/img8.jpg'),
    ('Товар 9', 900.0, 100, 'Большое описание товара', './images/img9.jpg');

-- Создание товаров в корзине для пользователя 1
INSERT INTO carts (
    user_id,
	product_id,
	product_count
)
VALUES
    (1, 1, 10),
    (1, 2, 20),
    (1, 3, 1),
    (1, 4, 5);

-- Создание заказа и товаров в заказе
WITH userInfo AS ( 
    -- Получение информации о пользователе
    SELECT name, email, id FROM users 
    WHERE id = 1
), orderInfo AS (
    -- Создание заказа
    INSERT INTO orders(
        fio,
        email,
        user_id
    )
    SELECT * FROM userInfo
    RETURNING id
), cartProducts AS (
    -- Получение продуктов в корзине 
    SELECT product_id, product_count FROM carts
    WHERE user_id = 1
)
-- Добавление товаров в заказ
INSERT INTO orders_products(
    product_id,
    product_count,
    order_id
)
SELECT product_id, product_count, (SELECT id FROM orderInfo) AS order_id FROM cartProducts;

DELETE FROM carts
    WHERE user_id = 1;