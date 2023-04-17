package migrations

import (
	"fmt"

	"github.com/ganiszulfa/concise/backend/internal/models/keys"
)

var InitMetadataSQLUp = `
INSERT INTO metadata (key, value)
VALUES 
	('site name', 'default'),
	('site tagline', 'default'),
	('user name', 'default'),
	('user password', '%s')
;
`

func (m *Migrators) initMetadata() {
	key := "InitMetadata"
	sql := fmt.Sprintf(InitMetadataSQLUp, keys.KEY_DEFAULT_VALUE)
	_, err := runMigration(m, sql, key)
	if err != nil {
		panic(err)
	}
}
