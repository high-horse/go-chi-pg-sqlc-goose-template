package models

import "server/sql/database"

type DBConfig struct {
	DB *database.Queries
}