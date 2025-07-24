CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          title TEXT NOT NULL,
                          description TEXT,
                          price NUMERIC(10, 2) NOT NULL CHECK (price > 0),
                          image_url TEXT,
                          quantity INT NOT NULL CHECK (quantity > 0),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                          category TEXT NOT NULL
);