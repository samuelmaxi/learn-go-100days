package main

type UserRepository interface {
	GetUserByID(id int) (string, error)
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUsername(id int) (string, error) {
	return s.repo.GetUserByID(id)
}

func main() {

}
