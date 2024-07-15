package files

import (
	"errors"
	"time"
)

var (
	ErrNameRequired  = errors.New("name is required")
	ErrTypeRequired  = errors.New("type is required")
	ErrPathRequired  = errors.New("path is required")
	ErrOwnerRequired = errors.New("owner is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func New(name string, folderID int64) (*File, error) {
	now := time.Now()

	f := &File{
		Name:       name,
		FolderID:   folderID,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	err := f.Validate()

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) Validate() error {
	if f.Name == "" {
		return ErrNameRequired
	}
	if f.OwnerID == 0 {
		return ErrOwnerRequired
	}
	if f.Type == "" {
		return ErrTypeRequired
	}
	if f.Path == "" {
		return ErrPathRequired
	}
	return nil
}
