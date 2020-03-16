package model

type User interface {
	ID() uint
	Name() string
}

type SystemUser struct {
	id   uint
	name string
}

type EmptyUser struct{}

func (e EmptyUser) ID() uint {
	return 0
}

func (e EmptyUser) Name() string {
	return ""
}

func NovoUsuario(id uint, name string) User {
	return &SystemUser{
		id:   id,
		name: name,
	}
}

func (u *SystemUser) ID() uint {
	return u.id
}

func (u *SystemUser) Name() string {
	return u.name
}
