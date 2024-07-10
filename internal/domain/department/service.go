package department

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(department *Department) error {
	return s.repo.Create(department)
}

func (s *service) GetAll() ([]Department, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id uint) (*Department, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(department *Department) error {
	return s.repo.Update(department)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
