package seeders

import "gorm.io/gorm"

type Seeder struct {
	db *gorm.DB
}

func NewDBSeeder(db *gorm.DB) Seeder {
	return Seeder{
		db: db,
	}
}

func (seed Seeder) GetAllSeeder() error {
	if err := adminSeeder(seed.db); err != nil {
		return err
	}

	if err := userSeeder(seed.db); err != nil {
		return err
	}

	return nil
}
