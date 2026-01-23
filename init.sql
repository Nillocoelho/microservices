-- Create databases idempotently (PostgreSQL)
SELECT 'CREATE DATABASE order_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'order_db')\gexec
SELECT 'CREATE DATABASE payment_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'payment_db')\gexec
SELECT 'CREATE DATABASE shipping_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'shipping_db')\gexec

-- Use order database
\connect order_db

-- Create tables for order database
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id),
    product_code VARCHAR(100) NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL
);

-- Table to store inventory items
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Insert sample products (you can add more as needed)
INSERT INTO products (product_code, name, price, quantity) VALUES
('PROD001', 'Product 1', 99.99, 100),
('PROD002', 'Product 2', 149.99, 50),
('PROD003', 'Product 3', 199.99, 75),
('PROD004', 'Product 4', 49.99, 200),
('PROD005', 'Product 5', 299.99, 30)
ON CONFLICT (product_code) DO NOTHING;

-- Use payment database
\connect payment_db

-- Create tables for payment database
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    user_id INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bills (
    id SERIAL PRIMARY KEY,
    payment_id INT NOT NULL REFERENCES payments(id),
    bill_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Use shipping database
\connect shipping_db

-- Create tables for shipping database
CREATE TABLE IF NOT EXISTS shipping (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    delivery_days INT NOT NULL,
    status VARCHAR(50) DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
