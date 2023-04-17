package migrations

import (
	"fmt"

	"github.com/ganiszulfa/concise/backend/internal/models/keys"

	"github.com/ganiszulfa/concise/backend/pkg/pwd"
	"github.com/sirupsen/logrus"
)

var ResetPasswordSQLUp = `
UPDATE metadata SET value = '%s'
WHERE key = '%s';
`

func (m *Migrators) resetPassword() {
	key := "ResetPassword"

	hash, err := pwd.Hash(keys.KEY_DEFAULT_VALUE)
	if err != nil {
		panic(err)
	}

	sql := fmt.Sprintf(ResetPasswordSQLUp, hash, keys.KEY_USER_PASSWORD)
	isRun, err := runMigration(m, sql, key)
	if err != nil {
		panic(err)
	}
	if isRun {
		logrus.Infof("Your user password is %s", keys.KEY_DEFAULT_VALUE)
	}
}
