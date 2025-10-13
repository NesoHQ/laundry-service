package db

import (
	"fmt"
	"log"

	"github.com/enghasib/laundry_service/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(cnf *config.Config) {

	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		cnf.DBUserName,
		cnf.DBPassword,
		cnf.DBName,
	)

	m, err := migrate.New("file://db/migrations", dbURL)

	if err != nil {
		fmt.Println("Migration error:", err)
		log.Fatal("Migration failed!", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migrations applied successfully!")
}
