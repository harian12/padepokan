package migrations

import (
	"gorm.io/gorm"
	"log"
	"padepokan/app/models"
)

type Migration struct {
	Table interface{}
}

func RegisterTable() []Migration {
	return []Migration{
		//daftarkan entity atau tabel migration disini
		{Table: models.Nasabah{}},
		{Table: models.Transaction{}},
	}
}

func DBMigrate(db *gorm.DB) {
	for _, tabel := range RegisterTable() {
		errMigrate := db.Debug().AutoMigrate(tabel.Table)
		if errMigrate != nil {
			log.Fatal(errMigrate)
		}
	}
}
