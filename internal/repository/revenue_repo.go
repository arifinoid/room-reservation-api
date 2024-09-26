package repository

import (
	"database/sql"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type RevenueRepository interface {
	GetRevenue() (domain.Revenue, error)
}

type revenueRepo struct {
	db *sql.DB
}

func NewRevenueRepo(db *sql.DB) RevenueRepository {
	return &revenueRepo{db: db}
}

func (r *revenueRepo) GetRevenue() (domain.Revenue, error) {
	var revenue domain.Revenue

	err := r.db.QueryRow("SELECT COALESCE(SUM(total), 0) FROM bookings WHERE payment_status = $1 AND reservation_date::date = CURRENT_DATE", "paid").Scan(&revenue.TodayRevenue)
	if err != nil {
		return domain.Revenue{}, err
	}

	err = r.db.QueryRow("SELECT COALESCE(SUM(total), 0) FROM bookings WHERE payment_status = $1", "paid").Scan(&revenue.OverallRevenue)
	if err != nil {
		return domain.Revenue{}, err
	}

	return revenue, nil
}
