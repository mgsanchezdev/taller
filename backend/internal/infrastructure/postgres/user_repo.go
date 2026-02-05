package postgres

import (
	"context"
	"database/sql"
	"time"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repositories.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, u *domain.User) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (id, workshop_id, email, password_hash, name, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		u.ID, u.WorkshopID, u.Email, u.PasswordHash, u.Name, u.Role, u.Active, u.CreatedAt, u.UpdatedAt)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, email, password_hash, name, role, active, created_at, updated_at
		FROM users WHERE id = $1`, id).Scan(
		&u.ID, &u.WorkshopID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepo) GetByEmail(ctx context.Context, workshopID, email string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, workshop_id, email, password_hash, name, role, active, created_at, updated_at
		FROM users WHERE workshop_id = $1 AND LOWER(email) = LOWER($2)`, workshopID, email).Scan(
		&u.ID, &u.WorkshopID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepo) ListByWorkshop(ctx context.Context, workshopID string) ([]*domain.User, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, workshop_id, email, password_hash, name, role, active, created_at, updated_at
		FROM users WHERE workshop_id = $1 ORDER BY name`, workshopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.WorkshopID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &u)
	}
	return list, nil
}

func (r *UserRepo) Update(ctx context.Context, u *domain.User) error {
	u.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE users SET email=$2, name=$3, role=$4, active=$5, updated_at=$6 WHERE id=$1`,
		u.ID, u.Email, u.Name, u.Role, u.Active, u.UpdatedAt)
	return err
}
