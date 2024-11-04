CREATE TABLE user_urls (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    url_id INTEGER REFERENCES urls(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, url_id)
);