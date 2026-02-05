package postgres

import (
	"context"
	"database/sql"
	"time"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"

	"github.com/google/uuid"
)

type VehicleRepo struct {
	db *sql.DB
}

func NewVehicleRepo(db *sql.DB) repositories.VehicleRepository {
	return &VehicleRepo{db: db}
}

func (r *VehicleRepo) Create(ctx context.Context, v *domain.Vehicle) error {
	v.ID = uuid.New().String()
	v.CreatedAt = time.Now()
	v.UpdatedAt = v.CreatedAt
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO vehicles (id, workshop_id, plate, brand, model, year, vin, color, owner_name, owner_phone, owner_email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		v.ID, v.WorkshopID, v.Plate, v.Brand, v.Model, v.Year, v.VIN, v.Color, v.OwnerName, v.OwnerPhone, v.OwnerEmail, v.CreatedAt, v.UpdatedAt)
	return err
}

func (r *VehicleRepo) GetByID(ctx context.Context, id string) (*domain.Vehicle, error) {
	var v domain.Vehicle
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, plate, brand, model, year, vin, color, owner_name, owner_phone, owner_email, created_at, updated_at
		FROM vehicles WHERE id = $1`, id).Scan(
		&v.ID, &v.WorkshopID, &v.Plate, &v.Brand, &v.Model, &v.Year, &v.VIN, &v.Color,
		&v.OwnerName, &v.OwnerPhone, &v.OwnerEmail, &v.CreatedAt, &v.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &v, err
}

func (r *VehicleRepo) GetByPlate(ctx context.Context, workshopID, plate string) (*domain.Vehicle, error) {
	var v domain.Vehicle
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, plate, brand, model, year, vin, color, owner_name, owner_phone, owner_email, created_at, updated_at
		FROM vehicles WHERE workshop_id = $1 AND LOWER(plate) = LOWER($2)`, workshopID, plate).Scan(
		&v.ID, &v.WorkshopID, &v.Plate, &v.Brand, &v.Model, &v.Year, &v.VIN, &v.Color,
		&v.OwnerName, &v.OwnerPhone, &v.OwnerEmail, &v.CreatedAt, &v.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &v, err
}

func (r *VehicleRepo) ListByWorkshop(ctx context.Context, workshopID string, limit, offset int) ([]*domain.Vehicle, int, error) {
	var total int
	_ = r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM vehicles WHERE workshop_id = $1`, workshopID).Scan(&total)
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, workshop_id, plate, brand, model, year, vin, color, owner_name, owner_phone, owner_email, created_at, updated_at
		FROM vehicles WHERE workshop_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, workshopID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var list []*domain.Vehicle
	for rows.Next() {
		var v domain.Vehicle
		if err := rows.Scan(&v.ID, &v.WorkshopID, &v.Plate, &v.Brand, &v.Model, &v.Year, &v.VIN, &v.Color,
			&v.OwnerName, &v.OwnerPhone, &v.OwnerEmail, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, 0, err
		}
		list = append(list, &v)
	}
	return list, total, nil
}

func (r *VehicleRepo) Update(ctx context.Context, v *domain.Vehicle) error {
	v.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE vehicles SET plate=$2, brand=$3, model=$4, year=$5, vin=$6, color=$7, owner_name=$8, owner_phone=$9, owner_email=$10, updated_at=$11 WHERE id=$1`,
		v.ID, v.Plate, v.Brand, v.Model, v.Year, v.VIN, v.Color, v.OwnerName, v.OwnerPhone, v.OwnerEmail, v.UpdatedAt)
	return err
}

func (r *VehicleRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM vehicles WHERE id = $1`, id)
	return err
}

