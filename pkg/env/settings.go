package env

type Settings struct {
	Message string
}

func (s *Settings) Init() error {
	s.Message = "Hello World"
	return nil
}
