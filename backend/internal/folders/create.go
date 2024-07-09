package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	f := new(Folder)

	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := Create(h.db, f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f.ID = id

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(f)
}

func Create(db *sql.DB, f *Folder) (int64, error) {
	stmt := `INSERT INTO folders (name, parent_id, created_at, modified_at) VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(stmt, f.Name, f.ParentID, f.CreatedAt, f.ModifiedAt)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
