package mock

import _ "embed"

//go:embed initdb.mock.sql
var InitdbMockSql string
