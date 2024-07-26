package files

import (
	"database/sql"
	"time"
)

func Restore(db *sql.DB, id int64) error {
	stmt := `UPDATE files SET deleted_at=NULL, modified_at=$1 WHERE id=$2`
	_, err := db.Exec(stmt, time.Now(), id)

	return err
}
