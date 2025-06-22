CREATE TABLE products (
                          product_id SERIAL PRIMARY KEY,
                          title TEXT NOT NULL,
                          description TEXT,
                          price NUMERIC(10,2),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


