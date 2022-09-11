-- Запросы

SELECT * FROM products WHERE active = true;

SELECT id FROM products WHERE name LIKE 'Товар%';

SELECT price FROM products WHERE id = 888;

SELECT name FROM users WHERE email = 'test1@test.com';