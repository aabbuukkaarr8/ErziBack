CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2),
    image_url TEXT,
    quantity int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


