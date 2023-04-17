package migrations

var CreateSessionTableSQLUp = `
	CREATE TABLE sessions (
		id varchar(100) PRIMARY KEY,
		created_at TIMESTAMPTZ,
		expired_at TIMESTAMPTZ,
		is_deleted boolean
	 );
`

func (m *Migrators) createSessionTable() {
	key := "CreateSessionTable"
	_, err := runMigration(m, CreateSessionTableSQLUp, key)
	if err != nil {
		panic(err)
	}
}
