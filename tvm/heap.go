package tvm

type heap struct {
	value []interface{}
}

func newHeap() *heap {
	return &heap{
		value: make([]interface{}, 0),
	}
}

func (h *heap) Add(object interface{}) {
	h.value = append(h.value, object)
}

func (h *heap) Get(index int64) interface{} {
	return h.value[index]
}

func (h *heap) GetAll() []interface{} {
	return h.value
}

func (h *heap) Swap(newHeap []interface{}) {
	h.value = newHeap
}

func (h *heap) SetObjectToPointer(pointer *pointer, object interface{}) {
	h.value[pointer.GetAddress()] = object
}

func (h *heap) GetObjectFromPointer(pointer *pointer) interface{} {
	return h.value[pointer.GetAddress()]

}
