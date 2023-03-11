package migrations

import (
	"gorm.io/gorm"
)

func prepare(db *gorm.DB) (prevMigrations map[string]bool, err error) {

	err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			key varchar(255) primary key
	 	);
	`).Error
	if err != nil {
		return
	}

	var prevMigrationSlices []Migration
	err = db.Find(&prevMigrationSlices).Error
	if err != nil {
		return
	}

	prevMigrations = make(map[string]bool, len(prevMigrationSlices))
	for _, migration := range prevMigrationSlices {
		prevMigrations[migration.Key] = true
	}

	return
}

func isExist(m *Migrators, key string) bool {
	if _, exist := m.prevMigrations[key]; !exist {
		return false
	}
	return true
}

func create(m *Migrators, key string) (err error) {
	migration := Migration{Key: key}
	err = m.db.Create(&migration).Error
	return
}

func runMigration(m *Migrators, sql, key string) (isRun bool, err error) {

	if isExist(m, key) {
		return
	}

	err = m.db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	if err := create(m, key); err != nil {
		panic(err)
	}

	isRun = true
	return
}
