package imitation

import (
	"fits-health/internal/entity"
)

type RepoUserMockTraditional struct {
	Register func(user entity.User) error
	GetAll   func() []entity.User
	GetID    func(id int) (entity.User, error)
	GetEmail func(email string) (entity.User, error)
	Update   func(id int, user entity.User) error
	Delete   func(id int) error
	// NewRepo  func(db *gorm.DB) domain.AdapterRepository
}

func (r *RepoUserMockTraditional) RegisterUser(user entity.User) error {
	return r.Register(user)
}

func (r *RepoUserMockTraditional) GetAllUsers() []entity.User {
	return r.GetAll()
}

func (r *RepoUserMockTraditional) GetUserByID(id int) (user entity.User, err error) {
	return r.GetID(id)
}

func (r *RepoUserMockTraditional) GetUserByEmail(email string) (entity.User, error) {
	return r.GetEmail(email)
}

func (r *RepoUserMockTraditional) UpdateUserByID(id int, user entity.User) error {
	return r.Update(id, user)
}

func (r *RepoUserMockTraditional) DeleteUserByID(id int) error {
	return r.Delete(id)
}

// func (r *RepoMockTraditional) NewRepository(db *gorm.DB) domain.AdapterRepository {
// 	return r.NewRepo(db)
// }
