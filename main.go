// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Page struct {
// 	data     []byte
// 	frame    int
// 	dirty    bool
// 	accessed bool
// }

// type FrameAllocator interface {
// 	AllocateFrame() int
// 	FreeFrame(int)
// }

// type MemoryAllocator struct {
// 	frames     []Page
// 	frameMap   map[int]Page
// 	framePool  sync.Pool
// 	frameCount int
// }

// func NewMemoryAllocator(frameCount int) *MemoryAllocator {
// 	return &MemoryAllocator{
// 		frames:     make([]Page, frameCount),
// 		frameMap:   make(map[int]Page),
// 		framePool:  sync.Pool{New: func() interface{} { return new(Page) }},
// 		frameCount: frameCount,
// 	}
// }

// func (ma *MemoryAllocator) AllocatePage() *Page {
// 	frame := ma.framePool.Get().(*Page)
// 	frame.data = make([]byte, 4096)
// 	frame.dirty = false
// 	frame.accessed = false
// 	ma.frames[frame.frame] = *frame
// 	ma.frameMap[frame.frame] = *frame
// 	return frame
// }

// func (ma *MemoryAllocator) FreePage(page *Page) {
// 	ma.frames[page.frame] = Page{}
// 	ma.frameMap[page.frame] = Page{}
// 	ma.framePool.Put(page)
// }

// type MMU struct {
// 	memoryAllocator *MemoryAllocator
// }

// func NewMMU(memoryAllocator *MemoryAllocator) *MMU {
// 	return &MMU{
// 		memoryAllocator: memoryAllocator,
// 	}
// }

// func (mmu *MMU) MapPage(virtualAddress int, physicalAddress int) {
// 	page := mmu.memoryAllocator.AllocatePage()
// 	page.frame = physicalAddress
// 	mmu.memoryAllocator.frames[physicalAddress] = *page
// }

// func (mmu *MMU) UnmapPage(virtualAddress int) {
// 	mmu.memoryAllocator.FreePage(&mmu.memoryAllocator.frames[virtualAddress])
// 	mmu.memoryAllocator.frames[virtualAddress] = Page{}
// }

// func (mmu *MMU) ReadBytes(virtualAddress int, offset ...int) []byte {
// 	page := mmu.memoryAllocator.frames[virtualAddress]
// 	return page.data[offset[0]:offset[1]]
// }

// func (mmu *MMU) WriteByte(virtualAddress int, datapos int, value byte) {
// 	page := mmu.memoryAllocator.frames[virtualAddress]
// 	fmt.Println(virtualAddress % 4096)
// 	page.data[datapos%4096] = value
// 	page.dirty = true
// }

// func main() {
// 	mmu := NewMMU(NewMemoryAllocator(1024))

// 	mmu.MapPage(0x01, 0x01)

// 	str := "Hello, world!"
// 	bytes := []byte(str)

// 	for i, b := range bytes {
// 		mmu.WriteByte(0x01, i, b)
// 	}

// 	fmt.Println(string(mmu.ReadBytes(0x01, 0, len(str))))

// }
package main

import (
	"mary_guica/pkg/tvm"
	"mary_guica/pkg/tvm/pkg/runtime"
)

func main() {
	str := []byte("oi me chamo igor")
	c := []byte{
		runtime.LOAD_STRING, 0x00, byte(len(str)),
	}

	c = append(c, str...)
	c = append(c, []byte{
		runtime.PRINT, 0x00,
	}...)
	c = append(c, runtime.S_THREAD, 0x0)
	c = append(c, runtime.HALT)

	vm := tvm.NewTVM(&tvm.ControlPlaneConfiguration{
		MemoryManager: tvm.MemoryManagerConfig{
			FrameSize: 1024,
		},
		ThreadManager: tvm.ThreadManagerConfig{},
		ProgramManager: tvm.ProgramManagerConfig{
			Code: c,
		},
	})

	vm.ExecuteCode(c)

	// tvm.T()
	// dsl, err := engine.File("main.todelovers")
	// if err != nil {
	// 	panic(err)
	// }
	// symbleTable := engine.NewSymbolTable()
	// lexer := engine.NewLexer(dsl).Tokenize()
	// nodeFactory := engine.NewNodeFactory()
	// assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(false)

	// root := assembler.GetRoot()
	// root.RegisterSymbols(symbleTable, nil)
	// engine.PrintSymbolTable(symbleTable)

	// engine.PrintByteCode(root.GenerateIntermediateCode(symbleTable))

	// 	ch := make(chan interface{})
	// 	run := r.NewRuntime(ch, control.NewControlPlane(&control.ControlPlaneConfiguration{
	// 		MemoryManager: control.MemoryManagerConfig{
	// 			FrameSize: 1024,
	// 		},
	// 		ProgramManager: control.ProgramManagerConfig{
	// 			Code: []byte{1, 2},
	// 		},
	// 	}))
	// 	env := e.NewEnvironment(ch)
	// 	go run.Requester()

	// 	go func() {
	// 		env.M(func(m r.MemoryManager) {
	// 			m.AllocateStack(1000)
	// 			stack := m.Stack()
	// 			stack.Push([]byte("hellllloooooow"))
	// 			stack.Push([]byte("wooooooords"))
	// 		})
	// 	}()
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		env.M(func(m r.MemoryManager) {
	// 			data, _ := m.Stack().Pop()

	// 			fmt.Println(string(data))
	// 		})

	// 	}

	// 	ch1 := make(chan int)
	// 	cho1 := make(chan e.Output)

	// 	ch2 := make(chan int)
	// 	cho2 := make(chan e.Output)

	// 	ch3 := make(chan int)
	// 	cho3 := make(chan e.Output)

	// 	go sendData(ch1, 2, 5)
	// 	go sendData(ch1, 4, 5)
	// 	go sendData(ch1, 7, 5)
	// 	go sendData(ch1, 6, 5)

	// 	go sendData(ch2, 2, 3)
	// 	go sendData(ch3, 5, 10)

	// 	r := runtime.NewNotifier()
	// 	r.RegisterWatcher(1, &runtime.RuntimeRunCommandWatcher{})

	// 	go func() {
	// 		for {
	// 			select {
	// 			case o := <-cho1:
	// 				r.NotifyWatchers(1, 1)
	// 				fmt.Println(o.Value)
	// 			case o := <-cho2:
	// 				fmt.Println(o.Value)
	// 			case o := <-cho3:
	// 				fmt.Println(o.Value)

	// 			}
	// 		}
	// 	}()

	// 	e.Teste(map[int]e.Input{
	// 		1: {
	// 			Id:  1,
	// 			In:  ch1,
	// 			Out: cho1,
	// 		},
	// 		2: {
	// 			Id:  2,
	// 			In:  ch2,
	// 			Out: cho2,
	// 		},
	// 		3: {
	// 			Id:  3,
	// 			In:  ch3,
	// 			Out: cho3,
	// 		},
	// 	})
	// }

	// func sendData(ch chan int, data int, s int64) {
	// 	for {
	// 		time.Sleep(time.Duration(s) * time.Second)
	// 		ch <- data
	// 	}

}
