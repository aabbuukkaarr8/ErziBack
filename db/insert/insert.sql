-- Вставка пользователей (id — UUID)
INSERT INTO users (id, username, email, password, role)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'Abu-bakr', 'abu@example.com',    'hashed_password_1', 'admin'),
    ('550e8400-e29b-41d4-a716-446655440001', 'Maga',     'maga@example.com',   'hashed_password_2', 'user'),
    ('550e8400-e29b-41d4-a716-446655440002', 'Baga',     'baga@example.com',   'hashed_password_3', 'user');

-- Вставка продуктов (остается без изменений)
INSERT INTO products (title, description, price, image_url, category, quantity)
VALUES
    ('Вода талая 10 литров', 'горная вода, одна из самых чистых в мире!', 250,  '', 'equipment', 20),
    ('Ачалуки 1,5 литра',     'Газированная вода из самых глубин',        100.99, '', 'honey-jam', 20),
    ('Куллер',                'Кулер для дома или офиса',                  8000,   '', 'equipment', 30);

-- Вставка корзин: теперь user_id — UUID
INSERT INTO carts (user_id, status)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'active'),
    ('550e8400-e29b-41d4-a716-446655440001', 'active'),
    ('550e8400-e29b-41d4-a716-446655440002', 'active');

-- Вставка товаров в корзины (cart_id у вас, возможно, всё ещё SERIAL)
INSERT INTO cart_items (cart_id, product_id, quantity)
VALUES
    (1, 1, 2),  -- Abu-bakr: 2×Product 1
    (1, 3, 1),  -- Abu-bakr: 1×Product 3
    (2, 2, 3),  -- Maga:     3×Product 2
    (3, 1, 1),  -- Baga:     1×Product 1
    (3, 2, 1);  -- Baga:     1×Product 2
