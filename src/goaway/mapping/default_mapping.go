package mapping

type DefaultMapping struct{}

func (m *DefaultMapping) Matches(uri string) bool {
	return true
}

func (m *DefaultMapping) TargetHost() string {
	return "localhost:9999"
}