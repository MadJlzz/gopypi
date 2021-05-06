package listing

type Distribution string

type Package struct {
	Name          string
	Version       string
	Distributions []Distribution
}
