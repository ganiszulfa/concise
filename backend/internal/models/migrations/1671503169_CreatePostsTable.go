package migrations

var CreatePostsTableSQLUp = `
	CREATE TABLE posts (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		title varchar(255) NOT NULL,
		slug varchar(255) NOT NULL,
		content text NOT NULL,
		published_at TIMESTAMPTZ,
		is_published boolean,
		is_page boolean,
		is_deleted boolean
	 );
`

func (m *Migrators) CreatePostsTable() {
	key := "CreatePostsTable"
	err := runMigration(m, CreatePostsTableSQLUp, key)
	if err != nil {
		panic(err)
	}
}
