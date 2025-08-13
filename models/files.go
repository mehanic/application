package models

import "server-application/database"

type File struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type files struct{}

var Files = new(files)

func (files) Create(name string) (*File, error) {
	f := &File{Name: name}
	var insertedID int
	tx := database.Postgre.MustBegin()
	tx.QueryRow("INSERT INTO files (name) VALUES ($1) RETURNING id", name).Scan(&insertedID)
	err := tx.Commit()
	f.ID = insertedID
	return f, err
}

func (files) List(name string) ([]*File, error) {
	f := []*File{}
	err := database.Postgre.Select(&f, "SELECT id, name FROM files")
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (files) ByName(name string) (*File, error) {
	f := File{}
	err := database.Postgre.Get(&f, "SELECT id, name FROM files WHERE name=$1;", name)
	if err != nil {
		return nil, err
	}
	return &f, nil
}
