CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,  -- Ulanyjy ady
    password_hash TEXT NOT NULL,           -- Şifrlenen parol
    created_at TIMESTAMP DEFAULT NOW()     -- Hasabyň döredilen wagty
);