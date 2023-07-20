package repository

import (
	"app/domain/entity"
	"context"
)

type FacultyRepository interface {
	GetFacultyByID(ctx context.Context, id int) (*entity.Faculty, error)
}
