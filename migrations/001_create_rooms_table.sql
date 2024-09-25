-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    feature JSONB,
    published BOOLEAN DEFAULT FALSE,
    availability INT,
    images TEXT[]
);

-- +migrate Down
DROP TABLE IF EXISTS rooms;
