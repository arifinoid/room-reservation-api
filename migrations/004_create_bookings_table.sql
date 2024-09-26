
-- +migrate Up
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES rooms(id) ON DELETE CASCADE,
    rateplan_id INT REFERENCES rateplans(id) ON DELETE CASCADE,
    calendar_id INT REFERENCES calendars(id) ON DELETE CASCADE,
    reservation_number VARCHAR(255) UNIQUE NOT NULL,
    reservation_date TIMESTAMPTZ DEFAULT NOW(),
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    country VARCHAR(100),
    total DECIMAL(10, 2) NOT NULL,
    payment_status VARCHAR(50) NOT NULL
);
CREATE INDEX idx_reservation_number ON bookings (reservation_number);

-- +migrate Down
DROP INDEX IF EXISTS idx_reservation_number;
DROP TABLE IF EXISTS bookings;
