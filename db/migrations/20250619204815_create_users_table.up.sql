-- 000001_create_users_table.up.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role TEXT DEFAULT 'product',
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
