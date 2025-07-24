CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id    UUID NOT NULL
                           REFERENCES users(id)
                               ON DELETE CASCADE,
                       status TEXT NOT NULL DEFAULT 'active'

);
