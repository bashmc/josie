package models

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
	Ext  string `json:"ext"`
	Url  string `json:"url"`
}

type FileStore interface {
	Save(context.Context, *File) error
	Update(context.Context, *File) error
	GetOne(context.Context) (File, error)
	GetMany(context.Context) ([]File, error)
	Delete(context.Context, string) error
}
