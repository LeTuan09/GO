package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetByUserRepo() string {

	return "UserRepo"
}