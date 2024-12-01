package migration

import (
	"embed"
)

//go:embed sql_files/*.sql
var MigrationsFS embed.FS
