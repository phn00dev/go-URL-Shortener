CREATE TABLE url_access_logs (
    id SERIAL PRIMARY KEY,
    url_id INTEGER REFERENCES urls(id) ON DELETE CASCADE,
    accessed_at TIMESTAMP DEFAULT NOW(),                 
    ip_address VARCHAR(45),                           
    user_agent TEXT                                     
);