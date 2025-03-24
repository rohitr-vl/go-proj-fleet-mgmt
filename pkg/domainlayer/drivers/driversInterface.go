package drivers

import (
	"context"
)

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, drivers Drivers) (*Drivers, error)
	All(ctx context.Context) ([]Drivers, error)
	GetByEmail(ctx context.Context, name string) (*Drivers, error)
	Update(ctx context.Context, id int64, updated Drivers) (*Drivers, error)
	Delete(ctx context.Context, id int64) error
}
