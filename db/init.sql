-- Criação da tabela products simplificada
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);

-- Inserção de alguns produtos de exemplo
INSERT INTO products (product_name, price) VALUES 
    ('Produto Teste 1', 29.99),
    ('Produto Teste 2', 49.99),
    ('Produto Teste 3', 19.99)
ON CONFLICT DO NOTHING;

-- Índices para melhor performance
CREATE INDEX IF NOT EXISTS idx_products_name ON products(product_name);
CREATE INDEX IF NOT EXISTS idx_products_price ON products(price);
