package files

import (
	"database/sql"
	"fmt"
)

func List(db *sql.DB, folderID int64) ([]*File, error) {
	stmt := `SELECT * FROM files where deleted_at is null and folder_id=$1`
	rows, err := db.Query(stmt, folderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*File
	for rows.Next() {
		f := new(File)
		err = rows.Scan(&f.ID, &f.FolderID, &f.OwnerID, &f.Name, &f.Type, &f.Path, &f.CreatedAt, &f.ModifiedAt, &f.DeletedAt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		files = append(files, f)
	}

	return files, nil
}
