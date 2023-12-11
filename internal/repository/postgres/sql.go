package postgres

const (
	queryInsertURLInfo = `
INSERT INTO url_info (url, alias, created_at) VALUES ($1, $2, NOW())
RETURNING id
`

	queryGetURLInfoByAlias = `
SELECT
	id,
	url,
	alias,
	created_at
FROM url_info
WHERE alias = $1
`
)
