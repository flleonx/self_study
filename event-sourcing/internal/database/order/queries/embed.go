package queries

import _ "embed"

//go:embed save_order.sql
var Save string

//go:embed list_event.sql
var ListEvent string
