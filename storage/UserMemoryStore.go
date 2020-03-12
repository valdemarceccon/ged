package storage

import "ged/model"

type UserMemoryStore struct {
	users []model.User
}

func NewUserMemoryStore() *UserMemoryStore {
	users := make([]model.User, 10)
	return &UserMemoryStore{users: users}
}

func (u *UserMemoryStore) Save(user model.User) {
	panic("implement me")
}

func (u *UserMemoryStore) Find(id uint) model.User {
	panic("implement me")
}

func (u *UserMemoryStore) Delete(user model.User) {
	panic("implement me")
}
