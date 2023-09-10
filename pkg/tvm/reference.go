package tvm

type Reference struct {
	value       int
	refCount    int
	isReachable bool
}
