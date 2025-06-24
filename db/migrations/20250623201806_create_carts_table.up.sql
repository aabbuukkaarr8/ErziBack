CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id INT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
