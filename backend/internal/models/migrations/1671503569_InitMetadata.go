package migrations

import (
	"fmt"

	"github.com/ganiszulfa/concise/backend/pkg/generate"
	"github.com/sirupsen/logrus"
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
	isRun, err := runMigration(m, sql, key)
	if err != nil {
		panic(err)
	}
	if isRun {
		logrus.Infof("Your user password is %s", pwd)
	}
}
