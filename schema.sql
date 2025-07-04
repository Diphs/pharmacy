CREATE DATABASE IF NOT EXISTS pharmacy_db;
USE pharmacy_db;

CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id VARCHAR(255) NOT NULL,
    medicine_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some test data
INSERT INTO transactions (transaction_id, medicine_name, quantity, price) VALUES
('TEST001', 'Aspirin', 2, 10.50),
('TEST002', 'Ibuprofen', 1, 8.25),
('TEST003', 'Paracetamol', 3, 12.00);
