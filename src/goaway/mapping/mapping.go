package mapping

type Mapping interface {
	Matches(uri string) bool
	TargetHost() string
}