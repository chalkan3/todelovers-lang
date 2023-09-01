package engine

type scope struct {
	Variables map[string]int
	Parent    *scope
}
