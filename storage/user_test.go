package storage

import (
	"ged/model"
	"testing"
)

func TestUserMemoryStore(t *testing.T) {

	store := NewUserMemoryStore()

	for _, u := range []model.User{
		model.NovoUsuario(0, "teste1"),
		model.NovoUsuario(0, "teste2"),
		model.NovoUsuario(0, "teste3"),
	} {
		store.Save(u)
	}

	if len(store.List()) != 3 {
		t.Error(" não tem 3")
	}

}
