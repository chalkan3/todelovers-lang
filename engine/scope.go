package engine

type Scope struct {
	Variables map[string]int
	Parent    *Scope
}
