package store

import "github.com/YerkanatSultanov/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
