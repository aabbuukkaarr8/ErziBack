CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id INT,
                       status TEXT NOT NULL DEFAULT 'active'

);
