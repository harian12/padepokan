package seeders

import (
	"fmt"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func ListSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		/*daftarkan seeder disini*/
		//{Seeder: UserSeeder(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range ListSeeders(db) {
		fmt.Println(seeder)
	}
	return nil
}
