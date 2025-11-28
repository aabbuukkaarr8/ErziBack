CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       username TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role TEXT DEFAULT 'user',
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);