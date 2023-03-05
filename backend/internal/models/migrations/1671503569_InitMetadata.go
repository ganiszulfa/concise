package migrations

import (
	"fmt"

	"github.com/ganiszulfa/concise/backend/pkg/generate"
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
	pwd := generate.RandAlphabetsLowerCase(24)
	sql := fmt.Sprintf(InitMetadataSQLUp, pwd)
	err := runMigration(m, sql, key)
	if err != nil {
		panic(err)
	}
}
