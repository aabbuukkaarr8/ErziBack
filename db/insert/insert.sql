INSERT INTO users (username, email, password, role)
VALUES
    ('Abu-bakr', 'abu@example.com', 'hashed_password_1', 'admin'),
    ('Maga', 'maga@example.com', 'hashed_password_2', 'user'),
    ('Baga', 'Baga@example.com', 'hashed_password_3', 'user');

-- Вставка продуктов
INSERT INTO products (title, description, price, image_url, category, quantity)
VALUES
    ('Вода талая 10 литров', 'горная вода, одна из самых чистых в мире!', 250, '', 'equipment', 20),
    ('Ачалуки 1,5 литра', 'Газированная вода из самых глубин', 100.99, '', 'honey-jam',20),
    ('Куллер', 'Куллер для дома или офиса', 8000, '', 'equipment', 30);

-- Вставка корзин
INSERT INTO carts (user_id, status)
VALUES
    (1, 'active'),
    (2, 'active'),
    (3, 'active');

-- Вставка товаров в корзины
INSERT INTO cart_items (cart_id, product_id, quantity)
VALUES
    (1, 1, 2),  -- Alice: 2 шт. Product A
    (1, 3, 1),  -- Alice: 1 шт. Product C
    (2, 2, 3),  -- Bob: 3 шт. Product B
    (3, 1, 1),  -- Carol: 1 шт. Product A
    (3, 2, 1);  -- Carol: 1 шт. Product B