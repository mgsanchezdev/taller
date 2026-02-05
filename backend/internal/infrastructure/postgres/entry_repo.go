package postgres

import (
	"context"
	"database/sql"
	"time"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"

	"github.com/google/uuid"
)

type EntryRepo struct {
	db *sql.DB
}

func NewEntryRepo(db *sql.DB) repositories.EntryRepository {
	return &EntryRepo{db: db}
}

func (r *EntryRepo) Create(ctx context.Context, e *domain.WorkshopEntry) error {
	e.ID = uuid.New().String()
	e.EntryDate = time.Now()
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
	if e.EntryNumber == "" {
		e.EntryNumber = "E-" + e.ID[:8]
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO workshop_entries (id, workshop_id, vehicle_id, entry_number, entry_date, mileage, fuel_level, notes, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		e.ID, e.WorkshopID, e.VehicleID, e.EntryNumber, e.EntryDate, e.Mileage, e.FuelLevel, e.Notes, e.Status, e.CreatedAt, e.UpdatedAt)
	return err
}

func (r *EntryRepo) GetByID(ctx context.Context, id string) (*domain.WorkshopEntry, error) {
	var e domain.WorkshopEntry
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, vehicle_id, entry_number, entry_date, mileage, fuel_level, notes, status, created_at, updated_at
		FROM workshop_entries WHERE id = $1`, id).Scan(
		&e.ID, &e.WorkshopID, &e.VehicleID, &e.EntryNumber, &e.EntryDate, &e.Mileage, &e.FuelLevel, &e.Notes, &e.Status, &e.CreatedAt, &e.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &e, err
}

func (r *EntryRepo) GetByEntryNumber(ctx context.Context, workshopID, entryNumber string) (*domain.WorkshopEntry, error) {
	var e domain.WorkshopEntry
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, vehicle_id, entry_number, entry_date, mileage, fuel_level, notes, status, created_at, updated_at
		FROM workshop_entries WHERE workshop_id = $1 AND entry_number = $2`, workshopID, entryNumber).Scan(
		&e.ID, &e.WorkshopID, &e.VehicleID, &e.EntryNumber, &e.EntryDate, &e.Mileage, &e.FuelLevel, &e.Notes, &e.Status, &e.CreatedAt, &e.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &e, err
}

func (r *EntryRepo) ListByWorkshop(ctx context.Context, workshopID string, limit, offset int) ([]*domain.WorkshopEntry, int, error) {
	var total int
	_ = r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM workshop_entries WHERE workshop_id = $1`, workshopID).Scan(&total)
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, workshop_id, vehicle_id, entry_number, entry_date, mileage, fuel_level, notes, status, created_at, updated_at
		FROM workshop_entries WHERE workshop_id = $1 ORDER BY entry_date DESC LIMIT $2 OFFSET $3`, workshopID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var list []*domain.WorkshopEntry
	for rows.Next() {
		var e domain.WorkshopEntry
		if err := rows.Scan(&e.ID, &e.WorkshopID, &e.VehicleID, &e.EntryNumber, &e.EntryDate, &e.Mileage, &e.FuelLevel, &e.Notes, &e.Status, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, 0, err
		}
		list = append(list, &e)
	}
	return list, total, nil
}

func (r *EntryRepo) ListByVehicle(ctx context.Context, vehicleID string) ([]*domain.WorkshopEntry, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, workshop_id, vehicle_id, entry_number, entry_date, mileage, fuel_level, notes, status, created_at, updated_at
		FROM workshop_entries WHERE vehicle_id = $1 ORDER BY entry_date DESC`, vehicleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*domain.WorkshopEntry
	for rows.Next() {
		var e domain.WorkshopEntry
		if err := rows.Scan(&e.ID, &e.WorkshopID, &e.VehicleID, &e.EntryNumber, &e.EntryDate, &e.Mileage, &e.FuelLevel, &e.Notes, &e.Status, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	return list, nil
}

func (r *EntryRepo) Update(ctx context.Context, e *domain.WorkshopEntry) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE workshop_entries SET mileage=$2, fuel_level=$3, notes=$4, status=$5, updated_at=$6 WHERE id=$1`,
		e.ID, e.Mileage, e.FuelLevel, e.Notes, e.Status, e.UpdatedAt)
	return err
}
