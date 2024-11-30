package postgres

import (
	"context"
	"errors"
	"fmt"
	"timeline/internal/repository/models"
)

var (
	ErrServiceNotFound = errors.New("service not found")
)

// Добавление организации предоставляемой услуги
func (p *PostgresRepo) ServiceAdd(ctx context.Context, service *models.Service) (int, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `
		INSERT INTO services
		(org_id, name, cost, description)
		VALUES($1, $2, $3, $4)
		RETURNING service_id;
	`
	var serviceID int
	if err := tx.QueryRowContext(ctx, query,
		service.OrgID,
		service.Name,
		service.Cost,
		service.Description,
	).Scan(&serviceID); err != nil {
		return 0, err
	}
	if tx.Commit() != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return serviceID, nil
}

// Обновление информации о предоставляемой услуге
func (p *PostgresRepo) ServiceUpdate(ctx context.Context, service *models.Service) error {
	tx, err := p.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `
		UPDATE services
		SET
			name = $1,
			cost = $2,
			description = $3
		WHERE service_id = $4 
		AND org_id = $5;
	`
	if err = tx.QueryRowContext(ctx, query,
		service.Name,
		service.Cost,
		service.Description,
		service.ServiceID,
		service.OrgID,
	).Err(); err != nil {
		return err
	}
	if tx.Commit() != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (p *PostgresRepo) Service(ctx context.Context, ServiceID, OrgID int) (*models.Service, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `
		SELECT service_id, org_id, name, cost, description
		FROM services
		WHERE service_id = $1
		AND org_id = $2;
	`
	var Service models.Service
	if err = tx.GetContext(ctx, &Service, query, &ServiceID, &OrgID); err != nil {
		return nil, err
	}
	if tx.Commit() != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return &Service, nil
}

// Получение списка услуг, предоставляемых организацией
func (p *PostgresRepo) ServiceList(ctx context.Context, OrgID int) ([]*models.Service, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `
		SELECT service_id, org_id, name, cost, description
		FROM services
		WHERE org_id = $1;
	`
	Services := make([]*models.Service, 0, 3)
	if err = tx.SelectContext(ctx, &Services, query, &OrgID); err != nil {
		return nil, err
	}
	if tx.Commit() != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return Services, nil
}

// Удаление услуги, предоставляемой организации и удаление связи с работниками
func (p *PostgresRepo) ServiceDelete(ctx context.Context, ServiceID, OrgID int) error {
	tx, err := p.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `
		DELETE 
		FROM services
		WHERE service_id = $1
		AND org_id = $2;
	`
	rows, err := tx.ExecContext(ctx, query, &ServiceID, &OrgID)
	if err != nil {
		return err
	}
	if rows != nil {
		if _, NoAffectedRows := rows.RowsAffected(); NoAffectedRows != nil {
			return ErrServiceNotFound
		}
	}
	if tx.Commit() != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (p *PostgresRepo) ServiceWorkerList(ctx context.Context, ServiceID, OrgID int) ([]*models.Worker, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := `SELECT worker_id, org_id, first_name, last_name, position, degree
		FROM workers
		WHERE worker_id IN (SELECT
								worker_id 
							FROM worker_services 
							WHERE service_id = $1)
		AND org_id = $2;
	`
	Workers := make([]*models.Worker, 0, 1)
	if err = tx.SelectContext(ctx, &Workers, query, &ServiceID, &OrgID); err != nil {
		return nil, err
	}
	if tx.Commit() != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return Workers, nil
}
