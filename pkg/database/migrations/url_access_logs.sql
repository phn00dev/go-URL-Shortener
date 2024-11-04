CREATE TABLE url_access_logs (
    id SERIAL PRIMARY KEY,
    url_id INTEGER REFERENCES urls(id) ON DELETE CASCADE, -- Baglanşykly URL
    accessed_at TIMESTAMP DEFAULT NOW(),                 -- Girilen wagt
    ip_address VARCHAR(45),                              -- IP Salgy
    user_agent TEXT                                      -- Brauzer maglumatlary
);