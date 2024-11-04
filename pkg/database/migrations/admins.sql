CREATE TYPE admin_role AS ENUM('super_admin','admin');

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,  
    password_hash  VARCHAR(255) NOT NULL,       
    admin_role  admin_role NOT NULL DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT NOW()   
);