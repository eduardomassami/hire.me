package domain

type service struct {
	urlRepo Repository
}

// initialize new URL
func NewURLService(urlRepo Repository) Service {
	return &service{urlRepo}
}

func (s *service) Get(alias string) ([]*URL, error) {
	return s.urlRepo.Get(alias)
}

func (s *service) GetMostUsed() ([]*URL, error) {
	return s.urlRepo.GetMostUsed()
}

func (s *service) SaveNoCustomAlias(user *URL) error {
	return s.urlRepo.Save(user)
}

func (s *service) SaveWithCustomAlias(user *URL) error {
	return s.urlRepo.Save(user)
}
