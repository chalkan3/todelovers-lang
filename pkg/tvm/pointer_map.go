package tvm

type pointerMap struct {
	value map[int64]*pointer
}

func newPointerMap() *pointerMap {
	return &pointerMap{
		value: make(map[int64]*pointer),
	}
}

func (pm *pointerMap) Set(key int64, pointer *pointer) {
	pm.value[key] = pointer
}

func (pm *pointerMap) Get(key int64) (*pointer, bool) {
	pp, ok := pm.value[key]
	return pp, ok
}
