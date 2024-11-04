CREATE TYPE admin_role AS ENUM('super_admin','admin');

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,  -- Ulanyjy ady
    email VARCHAR(100) UNIQUE NOT NULL,    -- Email
    password_hash TEXT NOT NULL,           -- Şifrlenen parol
    admin_role  admin_role NOT NULL DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT NOW()     -- Hasabyň döredilen wagty
);