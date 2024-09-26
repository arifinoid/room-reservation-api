# Room Reservation API

A RESTful API for managing a room reservation system. This API allows users to view, create, update, and delete rooms, rate plans, calendars, and bookings. It also provides features to manage availability and revenue tracking.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
  - [Local Installation](#local-installation)
  - [Running with Docker](#running-with-docker)
- [API Endpoints](#api-endpoints)
  - [Rooms](#rooms)
  - [Rate Plans](#rate-plans)
  - [Calendars](#calendars)
  - [Bookings](#bookings)
  - [Revenue](#revenue)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create, Read, Update, and Delete (CRUD) operations for:
  - Rooms
  - Rate Plans
  - Calendars
  - Bookings
- Manage availability in the calendar table when bookings are made or canceled.
- Retrieve total revenue for today and overall revenue filtered by payment status.
- Input validation for data integrity.

## Technologies

- Go (Golang)
- PostgreSQL
- Docker (for containerization)
- [sql-migrate](https://github.com/rubenv/sql-migrate) for database migrations
- [gorilla mux](https://github.com/gorilla/mux) for HTTP routing
- [validator](https://github.com/go-playground/validator) for input validation

## Installation

### Local Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/arifinoid/room-reservation-api.git
   cd room-reservation-api
   ```

2. Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

3. Set up your PostgreSQL database. Update the database connection settings in your code.

4. Start the API server:

   ```bash
   go run cmd/api/main.go
   ```

5. The API should be available at `http://localhost:8080`.

### Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t room-reservation-api .
   ```

2. Run the Docker container. Ensure to set environment variables for the database connection:

   ```bash
   docker run -d \
     --name room-reservation-api \
     -e DATABASE_URL="postgres://user:password@db:5432/room_reservation_db?sslmode=disable" \
     -p 8080:8080 \
     room-reservation-api
   ```

3. If you are using Docker Compose, you can run:

   ```bash
   docker-compose up
   ```

   The API should now be accessible at `http://localhost:8080`.

## API Endpoints

### Rooms

- **GET /api/v1/rooms**: Retrieve all rooms.
- **POST /api/v1/rooms**: Create a new room.
- **PUT /api/v1/rooms/{id}**: Update a room by ID.
- **DELETE /api/v1/rooms/{id}**: Delete a room by ID.

### Rate Plans

- **GET /api/v1/rateplans**: Retrieve all rate plans.
- **POST /api/v1/rateplans**: Create a new rate plan.
- **PUT /api/v1/rateplans/{id}**: Update a rate plan by ID.
- **DELETE /api/v1/rateplans/{id}**: Delete a rate plan by ID.

### Calendars

- **GET /api/v1/calendars**: Retrieve all calendars.
- **POST /api/v1/calendars**: Create a new calendar entry.
- **PUT /api/v1/calendars/{id}**: Update a calendar entry by ID.
- **DELETE /api/v1/calendars/{id}**: Delete a calendar entry by ID.

### Bookings

- **GET /api/v1/bookings**: Retrieve all bookings (with filtering options).
- **POST /api/v1/bookings**: Create a new booking.
- **PUT /api/v1/bookings/{id}**: Update a booking by ID.
- **DELETE /api/v1/bookings/{id}**: Delete a booking by ID.

### Revenue

- **GET /api/v1/revenue**: Retrieve today's and overall revenue filtered by payment status.

## Testing

You can use tools like `curl` or Postman to test the API endpoints. Ensure to provide the necessary headers and payloads as per the endpoint requirements.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
