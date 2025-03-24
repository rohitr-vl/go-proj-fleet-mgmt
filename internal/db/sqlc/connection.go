package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitiateDbConnection() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres password=super369 dbname=fleet_track_mgmt sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	return nil
}