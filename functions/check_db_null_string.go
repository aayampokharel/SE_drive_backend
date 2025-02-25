package functions

import (
	"database/sql"
)

func CheckDbNullString(dbNullStringType *sql.NullString) string {
	if dbNullStringType.Valid {
		return dbNullStringType.String

	} else {
		return ""
	}
}
