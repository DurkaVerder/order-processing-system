CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,  -- Хэшированный пароль
    user_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,  -- Связь с пользователем
    total_amount DECIMAL(10, 2) NOT NULL,               -- Общая сумма заказа
    status VARCHAR(50) NOT NULL,                        -- Текущий статус заказа
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,  -- Связь с заказом
    product_name VARCHAR(255) NOT NULL,                   -- Название товара/услуги
    quantity INT NOT NULL,                                -- Количество
    price DECIMAL(10, 2) NOT NULL,                        -- Цена за единицу
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_status_history (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,  -- Связь с заказом
    status VARCHAR(50) NOT NULL,                          -- Новый статус
    changed_by INT REFERENCES users(id) ON DELETE SET NULL,  -- Кто изменил статус
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,  -- Связь с заказом
    amount DECIMAL(10, 2) NOT NULL,                       -- Сумма платежа
    payment_method VARCHAR(50) NOT NULL,                  -- Способ оплаты (например, Stripe, PayPal)
    status VARCHAR(50) NOT NULL,                          -- Статус платежа (например, "pending", "completed", "failed")
    transaction_id VARCHAR(255),                          -- ID транзакции в платежной системе
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);