package storage

import "ged/model"

type FileStore struct {
	filename string
	owner    model.User
}
