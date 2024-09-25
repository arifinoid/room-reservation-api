-- +migrate Up
CREATE TABLE IF NOT EXISTS rateplans (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES rooms(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    detail TEXT,
    price DECIMAL(10, 2) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS rateplans;
