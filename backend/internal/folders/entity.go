package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequired = errors.New("name is required")
)

type Folder struct {
	ID         int64     `json:"id"`
	ParentID   int64     `json:"parent_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type FolderContent struct {
	Folder  *Folder          `json:"folder"`
	Content []FolderResource `json:"content"`
}

type FolderResource struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func New(name string, parentID int64) (*Folder, error) {
	now := time.Now()

	f := &Folder{
		Name:       name,
		ParentID:   parentID,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	err := f.Validate()

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Folder) Validate() error {
	if f.Name == "" {
		return ErrNameRequired
	}
	return nil
}
