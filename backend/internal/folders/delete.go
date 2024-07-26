package folders

import (
	"database/sql"
	"encoding/json"
	"my-dropbox/internal/files"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = deleteFiles(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: list folders

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(nil)
}

func deleteFiles(db *sql.DB, id int64) error {
	f, err := files.List(db, int64(id))
	if err != nil {
		return err
	}

	removedFiles := make([]*files.File, 0, len(f))
	for _, file := range f {
		err = files.Delete(db, file.ID)
		if err != nil {
			break
		}

		removedFiles = append(removedFiles, file)
	}

	if len(removedFiles) != len(f) {
		for _, file := range removedFiles {
			err = files.Restore(db, file.ID)
		}

		return err
	}

	return nil
}

func Delete(db *sql.DB, id int64) error {
	stmt := `UPDATE folders SET deleted_at=$1 WHERE id=$2`
	_, err := db.Exec(stmt, time.Now(), id)

	return err
}
