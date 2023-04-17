package migrations

import (
	"fmt"

	"github.com/ganiszulfa/concise/backend/internal/models/keys"
)

var InitMetadata2SQLUp = `
INSERT INTO metadata (key, value)
VALUES 
	('%s', '%s'),
	('%s', '%s')
;
`

func (m *Migrators) initMetadata2() {
	key := "InitMetadata2"
	sql := fmt.Sprintf(InitMetadata2SQLUp,
		keys.KEY_SITE_DESCRIPTION, keys.KEY_DEFAULT_VALUE,
		keys.KEY_ADMIN_ID, keys.KEY_DEFAULT_VALUE,
	)
	_, err := runMigration(m, sql, key)
	if err != nil {
		panic(err)
	}
}
