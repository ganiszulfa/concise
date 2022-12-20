package migrations

import (
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"gorm.io/gorm"
)

type Migrators struct {
	db             *gorm.DB
	prevMigrations map[string]bool
}

func Migrate(db *gorm.DB) (err error) {
	trace.Func()

	prevMigrations, err := prepare(db)
	if err != nil {
		return
	}

	migrator := Migrators{
		db:             db,
		prevMigrations: prevMigrations,
	}

	// PUT THE MIGRATION SCRIPTS HERE
	migrator.CreateMetadataTable()
	migrator.CreatePostsTable()

	return
}
