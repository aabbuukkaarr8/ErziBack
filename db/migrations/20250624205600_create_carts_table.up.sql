CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id INT
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
