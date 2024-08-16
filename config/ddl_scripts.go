package config

func DDLScript() string {
	return `
	CREATE TABLE IF NOT EXISTS "Student" (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		created_at TIMESTAMPTZ NOT NULL,
		modified_at TIMESTAMPTZ NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL
	);
	`
}