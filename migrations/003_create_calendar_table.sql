-- +migrate Up
CREATE TABLE IF NOT EXISTS calendars (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    rateplan_id INTEGER NOT NULL REFERENCES rateplans(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    availability INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS calendars;
