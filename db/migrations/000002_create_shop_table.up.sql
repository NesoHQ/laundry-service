
CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    unique_id UUID UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    location TEXT NOT NULL,
    contact VARCHAR(50) NOT NULL,
    payment_details TEXT,
    created_by UUID NOT NULL REFERENCES users(unique_id) ON DELETE CASCADE,
    shop_owner UUID NOT NULL REFERENCES users(unique_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
