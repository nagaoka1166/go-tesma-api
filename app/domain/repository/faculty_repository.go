package repository

import (
	"github.com/nagaoka166/go-tesma-api/app/domain/entity"
	"context"
)

type FacultyRepository interface {
	GetFacultyByID(ctx context.Context, id int) (*entity.Faculty, error)
}
