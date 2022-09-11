package main

import (
	"github.com/nguyenvulong/go-bank-backend/api"
	"github.com/nguyenvulong/go-bank-backend/migrations"
)

func main() {
	migrations.Migrate()
	api.StartApi()
}
