CREATE TABLE migrations (
    id SERIAL PRIMARY KEY,
    version VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);