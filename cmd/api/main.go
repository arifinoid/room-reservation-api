package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/config"
	"github.com/arifinoid/room-reservation-api/internal/database"
	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/repository"
	"github.com/arifinoid/room-reservation-api/internal/routes"
	"github.com/arifinoid/room-reservation-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

func runMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

var validate *validator.Validate

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	defer db.Close()

	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	validate = validator.New()
	validate.RegisterValidation("slug", lib.ValidateSlug)

	roomRepo := repository.NewRoomRepo(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService, validate)

	rateplanRepo := repository.NewRatePlanRepo(db)
	rateplanService := service.NewRatePlanService(rateplanRepo)
	rateplanHandler := handler.NewRatePlanHandler(rateplanService, validate)

	calendarRepo := repository.NewCalendarRepo(db)
	calendarService := service.NewCalendarService(calendarRepo)
	calendarHandler := handler.NewCandendarHandler(calendarService, validate)

	bookingRepo := repository.NewBookingRepo(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingHandler := handler.NewBookingHandler(bookingService, validate)

	revenueRepo := repository.NewRevenueRepo(db)
	revenueService := service.NewRevenueService(revenueRepo)
	revenueHandler := handler.NewRevenueHandler(revenueService)

	router := mux.NewRouter()
	routes.RegisterRoomRoutes(router, roomHandler)
	routes.RegisterRatePlanRoutes(router, rateplanHandler)
	routes.RegisterCalendarRoutes(router, calendarHandler)
	routes.RegisterBookingRoutes(router, bookingHandler)
	routes.RegisterRevenueRoutes(router, revenueHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
