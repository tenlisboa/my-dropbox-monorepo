package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := GetFolder(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	c, err := GetFolderContent(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fc := FolderContent{
		Folder:  f,
		Content: c,
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)
}

func GetFolder(db *sql.DB, id int64) (*Folder, error) {
	stmt := `SELECT * FROM folders WHERE id=$1`
	row := db.QueryRow(stmt, id)
	f := new(Folder)

	err := row.Scan(&f.ID, &f.Name, &f.ParentID, &f.CreatedAt, &f.ModifiedAt, &f.DeletedAt)

	return f, err
}

func getSubFolder(db *sql.DB, folderID int64) ([]*Folder, error) {
	stmt := `SELECT * FROM folders WHERE parent_id=$1 and deleted_at is null`
	row, err := db.Query(stmt, folderID)

	f := make([]*Folder, 0)

	for row.Next() {
		folder := new(Folder)
		err = row.Scan(&folder.ID, &folder.Name, &folder.ParentID, &folder.CreatedAt, &folder.ModifiedAt, &folder.DeletedAt)
		if err != nil {
			continue
		}
		f = append(f, folder)
	}

	return f, err
}

func GetFolderContent(db *sql.DB, id int64) ([]FolderResource, error) {
	subfolders, err := getSubFolder(db, id)

	if err != nil {
		return nil, err
	}

	fr := make([]FolderResource, 0, len(subfolders))
	for _, sf := range subfolders {
		fr = append(fr, FolderResource{
			ID:         sf.ID,
			Name:       sf.Name,
			Type:       "directory",
			ModifiedAt: sf.ModifiedAt,
			CreatedAt:  sf.CreatedAt,
		})
	}

	return fr, nil
}
