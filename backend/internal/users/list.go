package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *handler) ListAll(rw http.ResponseWriter, r *http.Request) {
	users, err := SelectAll(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(users)
}

func SelectAll(db *sql.DB) ([]*User, error) {
	stmt := `SELECT * FROM users where deleted_at is null`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.DeletedAt, &u.LastLoginAt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}
