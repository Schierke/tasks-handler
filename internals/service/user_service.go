package service

type UserRepository interface {
}

type userserviceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userserviceImpl {
	return &userserviceImpl{
		repo: repo,
	}
}
