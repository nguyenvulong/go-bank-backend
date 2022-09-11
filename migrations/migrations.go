package migrations

import (
	"github.com/rs/zerolog/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nguyenvulong/go-bank-backend/helpers"
	. "github.com/nguyenvulong/go-bank-backend/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()
	users := [2]User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	log.Info().Msg("Users created successfully!")
	// defer db.Close()
}

func Migrate() {
	db := helpers.ConnectDB()
	db.AutoMigrate(&User{}, &Account{})
	//defer db.Close()

	createAccounts()
}
