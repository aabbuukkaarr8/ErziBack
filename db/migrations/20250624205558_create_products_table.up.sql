CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          title TEXT NOT NULL,
                          description TEXT,
                          price NUMERIC(10, 2) NOT NULL CHECK (price > 0),
                          image_url TEXT,
                          is_active BOOL DEFAULT true,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                          category TEXT NOT NULL,
                          bulk_discount_quantity INT DEFAULT 0,
                          bulk_discount_price    NUMERIC(10, 2) DEFAULT 0
);