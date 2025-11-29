CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          title TEXT NOT NULL,
                          description TEXT,
                          image_url TEXT,
                          is_active BOOL DEFAULT true,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                          category TEXT NOT NULL,
                          prices JSONB DEFAULT '[]'::jsonb NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_products_prices ON products USING GIN (prices);