package tvm

type garbageCollector struct {
	tvm          *TVM
	aliveObjects []interface{}
}

func NewGarbageCollector(vm *TVM) *garbageCollector {
	return &garbageCollector{
		tvm:          vm,
		aliveObjects: make([]interface{}, 0),
	}
}

// Marca um objeto como vivo.
func (gc *garbageCollector) MarkObject(object interface{}) {
	gc.aliveObjects = append(gc.aliveObjects, object)
}

// Marca todos os objetos referenciados por um objeto.
func (gc *garbageCollector) MarkObjectReferences(object interface{}) {
	for _, reference := range gc.tvm.GetObjectReferences(object) {
		gc.MarkObject(reference)
	}
}

// Coleta o lixo.
func (gc *garbageCollector) CollectGarbage() {
	heap := gc.tvm.heap.GetAll()
	// Marca todos os objetos vivos.
	for _, object := range heap {
		gc.MarkObject(object)
	}

	// Remove todos os objetos não vivos do heap.
	for index, object := range heap {
		if !gc.IsObjectAlive(object) {
			gc.tvm.heap.Swap(heap[:index-1])
		}
	}
}

// Verifica se um objeto está vivo.
func (gc *garbageCollector) IsObjectAlive(object interface{}) bool {
	for _, aliveObject := range gc.aliveObjects {
		if aliveObject == object {
			return true
		}
	}

	return false
}
