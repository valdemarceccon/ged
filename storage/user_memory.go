package storage

import "ged/model"

type UserMemoryStore struct {
	users map[uint]model.User
}

func NewUserMemoryStore() *UserMemoryStore {
	users := make(map[uint]model.User, 0)
	return &UserMemoryStore{users: users}
}

func (u *UserMemoryStore) Save(user model.User) {
	if user.ID() == 0 {
		create(u, user)
	} else {
		updateUser(u, user)
	}
}

func create(u *UserMemoryStore, user model.User) {
	var gratesdId uint = 0

	for id := range u.users {
		if id > gratesdId {
			gratesdId = id
		}
	}

	u.users[gratesdId+1] = model.NovoUsuario(gratesdId+1, user.Name())
}

func updateUser(ums *UserMemoryStore, user model.User) {
	ums.users[user.ID()] = user
}

func (u *UserMemoryStore) Find(wantedId uint) model.User {
	for id, user := range u.users {
		if wantedId == id {
			return user
		}
	}

	return model.EmptyUser{}
}

func (u *UserMemoryStore) Delete(user model.User) {
	panic("implement me")
}

func (u *UserMemoryStore) List() []model.User {
	result := make([]model.User, 0, len(u.users))

	for _, v := range u.users {
		result = append(result, v)
	}

	return result
}
