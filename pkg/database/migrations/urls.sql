CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,         -- Uzyn URL (başlangyç salgysy)
    short_url VARCHAR(10) UNIQUE,      -- Gysga URL üçin unikal kod
    created_at TIMESTAMP DEFAULT NOW(), -- URL-iň döredilen wagty
    expires_at TIMESTAMP                -- URL-iň möhleti (islege görä)
);
