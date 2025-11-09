CREATE TABLE product_attributes (
                                    id SERIAL PRIMARY KEY,
                                    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                                    key TEXT NOT NULL,
                                    value TEXT NOT NULL,
                                    UNIQUE (product_id, key)
);

CREATE INDEX idx_product_attributes_product_id ON product_attributes(product_id);
CREATE INDEX idx_product_attributes_key ON product_attributes(key);
CREATE INDEX idx_product_attributes_key_value ON product_attributes(key, value);
