package migrations

var CreateMetadataTableSQLUp = `
	CREATE TABLE metadata (
		id SERIAL PRIMARY KEY,
		key varchar(255) UNIQUE NOT NULL,
		value text NOT NULL
	 );
`

func (m *Migrators) createMetadataTable() {
	key := "CreateMetadataTable"
	_, err := runMigration(m, CreateMetadataTableSQLUp, key)
	if err != nil {
		panic(err)
	}
}
