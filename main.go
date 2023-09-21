// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // )

// // type Page struct {
// // 	data     []byte
// // 	frame    int
// // 	dirty    bool
// // 	accessed bool
// // }

// // type FrameAllocator interface {
// // 	AllocateFrame() int
// // 	FreeFrame(int)
// // }

// // type MemoryAllocator struct {
// // 	frames     []Page
// // 	frameMap   map[int]Page
// // 	framePool  sync.Pool
// // 	frameCount int
// // }

// // func NewMemoryAllocator(frameCount int) *MemoryAllocator {
// // 	return &MemoryAllocator{
// // 		frames:     make([]Page, frameCount),
// // 		frameMap:   make(map[int]Page),
// // 		framePool:  sync.Pool{New: func() interface{} { return new(Page) }},
// // 		frameCount: frameCount,
// // 	}
// // }

// // func (ma *MemoryAllocator) AllocatePage() *Page {
// // 	frame := ma.framePool.Get().(*Page)
// // 	frame.data = make([]byte, 4096)
// // 	frame.dirty = false
// // 	frame.accessed = false
// // 	ma.frames[frame.frame] = *frame
// // 	ma.frameMap[frame.frame] = *frame
// // 	return frame
// // }

// // func (ma *MemoryAllocator) FreePage(page *Page) {
// // 	ma.frames[page.frame] = Page{}
// // 	ma.frameMap[page.frame] = Page{}
// // 	ma.framePool.Put(page)
// // }

// // type MMU struct {
// // 	memoryAllocator *MemoryAllocator
// // }

// // func NewMMU(memoryAllocator *MemoryAllocator) *MMU {
// // 	return &MMU{
// // 		memoryAllocator: memoryAllocator,
// // 	}
// // }

// // func (mmu *MMU) MapPage(virtualAddress int, physicalAddress int) {
// // 	page := mmu.memoryAllocator.AllocatePage()
// // 	page.frame = physicalAddress
// // 	mmu.memoryAllocator.frames[physicalAddress] = *page
// // }

// // func (mmu *MMU) UnmapPage(virtualAddress int) {
// // 	mmu.memoryAllocator.FreePage(&mmu.memoryAllocator.frames[virtualAddress])
// // 	mmu.memoryAllocator.frames[virtualAddress] = Page{}
// // }

// // func (mmu *MMU) ReadBytes(virtualAddress int, offset ...int) []byte {
// // 	page := mmu.memoryAllocator.frames[virtualAddress]
// // 	return page.data[offset[0]:offset[1]]
// // }

// // func (mmu *MMU) WriteByte(virtualAddress int, datapos int, value byte) {
// // 	page := mmu.memoryAllocator.frames[virtualAddress]
// // 	fmt.Println(virtualAddress % 4096)
// // 	page.data[datapos%4096] = value
// // 	page.dirty = true
// // }

// // func main() {
// // 	mmu := NewMMU(NewMemoryAllocator(1024))

// // 	mmu.MapPage(0x01, 0x01)

// // 	str := "Hello, world!"
// // 	bytes := []byte(str)

// // 	for i, b := range bytes {
// // 		mmu.WriteByte(0x01, i, b)
// // 	}

// // 	fmt.Println(string(mmu.ReadBytes(0x01, 0, len(str))))

// // }
package main

import (
	"mary_guica/pkg/tvm"
	"mary_guica/pkg/tvm/pkg/runtime"
)

func main() {
	str := []byte("Thread0")
	str2 := []byte("Thread0-1 ")
	str3 := []byte("Thread2 ")
	str4 := []byte("Thread3 ")

	c := []byte{
		runtime.LOAD_STRING, 0x00, byte(len(str)),
	}
	c2 := []byte{
		runtime.LOAD_STRING, 0x00, byte(len(str2)),
	}

	c3 := []byte{
		runtime.LOAD_STRING, 0x00, byte(len(str3)),
	}

	c4 := []byte{
		runtime.LOAD_STRING, 0x00, byte(len(str4)),
	}

	k := []byte{}
	k = append(k, c4...)
	k = append(k, str4...)
	k = append(k, runtime.PRINT, 0x0)

	t := []byte{}
	t = append(t, runtime.S_THREAD, byte(len(k)))
	t = append(t, k...)
	t = append(t, runtime.ST_THREAD)
	t = append(t, runtime.W_THREAD)
	t = append(t, c3...)
	t = append(t, str3...)
	t = append(t, runtime.PRINT, 0x0)

	c = append(c, str...)
	c = append(c, []byte{
		runtime.PRINT, 0x00,
	}...)
	c = append(c, runtime.S_THREAD, byte(len(t)))
	c = append(c, t...)
	c = append(c, runtime.ST_THREAD)
	c = append(c, runtime.W_THREAD)
	c = append(c, c2...)
	c = append(c, str2...)
	c = append(c, runtime.PRINT, 0x00)
	c = append(c, runtime.HALT)

	vm := tvm.NewTVM(&tvm.ControlPlaneConfiguration{
		MemoryManager: tvm.MemoryManagerConfig{
			FrameSize: 1024,
		},
		ThreadManager: tvm.ThreadManagerConfig{},
		ProgramManager: tvm.ProgramManagerConfig{
			Code: c,
		},
	}).Startup()
	vm.ExecuteCode(c)

	// 	// tvm.T()
	// 	// dsl, err := engine.File("main.todelovers")
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }
	// 	// symbleTable := engine.NewSymbolTable()
	// 	// lexer := engine.NewLexer(dsl).Tokenize()
	// 	// nodeFactory := engine.NewNodeFactory()
	// 	// assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(false)

	// 	// root := assembler.GetRoot()
	// 	// root.RegisterSymbols(symbleTable, nil)
	// 	// engine.PrintSymbolTable(symbleTable)

	// 	// engine.PrintByteCode(root.GenerateIntermediateCode(symbleTable))

	// 	// 	ch := make(chan interface{})
	// 	// 	run := r.NewRuntime(ch, control.NewControlPlane(&control.ControlPlaneConfiguration{
	// 	// 		MemoryManager: control.MemoryManagerConfig{
	// 	// 			FrameSize: 1024,
	// 	// 		},
	// 	// 		ProgramManager: control.ProgramManagerConfig{
	// 	// 			Code: []byte{1, 2},
	// 	// 		},
	// 	// 	}))
	// 	// 	env := e.NewEnvironment(ch)
	// 	// 	go run.Requester()

	// 	// 	go func() {
	// 	// 		env.M(func(m r.MemoryManager) {
	// 	// 			m.AllocateStack(1000)
	// 	// 			stack := m.Stack()
	// 	// 			stack.Push([]byte("hellllloooooow"))
	// 	// 			stack.Push([]byte("wooooooords"))
	// 	// 		})
	// 	// 	}()
	// 	// 	for {
	// 	// 		time.Sleep(5 * time.Second)
	// 	// 		env.M(func(m r.MemoryManager) {
	// 	// 			data, _ := m.Stack().Pop()

	// 	// 			fmt.Println(string(data))
	// 	// 		})

	// 	// 	}

	// 	// 	ch1 := make(chan int)
	// 	// 	cho1 := make(chan e.Output)

	// 	// 	ch2 := make(chan int)
	// 	// 	cho2 := make(chan e.Output)

	// 	// 	ch3 := make(chan int)
	// 	// 	cho3 := make(chan e.Output)

	// 	// 	go sendData(ch1, 2, 5)
	// 	// 	go sendData(ch1, 4, 5)
	// 	// 	go sendData(ch1, 7, 5)
	// 	// 	go sendData(ch1, 6, 5)

	// 	// 	go sendData(ch2, 2, 3)
	// 	// 	go sendData(ch3, 5, 10)

	// 	// 	r := runtime.NewNotifier()
	// 	// 	r.RegisterWatcher(1, &runtime.RuntimeRunCommandWatcher{})

	// 	// 	go func() {
	// 	// 		for {
	// 	// 			select {
	// 	// 			case o := <-cho1:
	// 	// 				r.NotifyWatchers(1, 1)
	// 	// 				fmt.Println(o.Value)
	// 	// 			case o := <-cho2:
	// 	// 				fmt.Println(o.Value)
	// 	// 			case o := <-cho3:
	// 	// 				fmt.Println(o.Value)

	// 	// 			}
	// 	// 		}
	// 	// 	}()

	// 	// 	e.Teste(map[int]e.Input{
	// 	// 		1: {
	// 	// 			Id:  1,
	// 	// 			In:  ch1,
	// 	// 			Out: cho1,
	// 	// 		},
	// 	// 		2: {
	// 	// 			Id:  2,
	// 	// 			In:  ch2,
	// 	// 			Out: cho2,
	// 	// 		},
	// 	// 		3: {
	// 	// 			Id:  3,
	// 	// 			In:  ch3,
	// 	// 			Out: cho3,
	// 	// 		},
	// 	// 	})
	// 	// }

	// 	// func sendData(ch chan int, data int, s int64) {
	// 	// 	for {
	// 	// 		time.Sleep(time.Duration(s) * time.Second)
	// 	// 		ch <- data
	// 	// 	}

	// }
	// package main

	// import (
	// 	"mary_guica/pkg/nando"
	// 	"mary_guica/pkg/tvm"

	// 	"time"
	// )

	// func main() {
	// w := wal.NewTLWAL()
	// w.CreateLogFile()

	// example := nando.NewHandler("example", func(r *nando.Request) (*nando.Response, error) {

	// 	record := &wal.Record{
	// 		Operation: "INSERT",
	// 		Table:     "jubileu",
	// 		Data: &wal.Data{
	// 			Key:   "1",
	// 			Value: "inserted",
	// 		},
	// 	}

	// 	w.Write(record, true)
	// 	return nil, nil
	// })

	// read := nando.NewHandler("read", func(r *nando.Request) (*nando.Response, error) {
	// 	records, err := w.ReadAll()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	newR := wal.Table(r.Data.(string), records)
	// 	for i, record := range newR {
	// 		fmt.Printf("Record %d:\n", i+1)
	// 		fmt.Printf("  Timestamp: %s\n", record.Timestamp)
	// 		fmt.Printf("  Operation: %s\n", record.Operation)
	// 		fmt.Printf("  Table: %s\n", record.Table)

	// 		fmt.Printf("  Data: %v\n", record.Data)
	// 	}
	// 	return nil, nil
	// })

	// teste1 := nando.NewHandler("teste1", func(r *nando.Request) (*nando.Response, error) {
	// 	record := &wal.Record{
	// 		Operation: "INSERT",
	// 		Table:     "kakitipiu",
	// 		Data: &wal.Data{
	// 			Key:   "1",
	// 			Value: "inserted",
	// 		},
	// 	}

	// 	w.Write(record, true)
	// 	return nil, nil
	// })

	// c := &nando.Client{}

	// go func() {
	// 	req := nando.NewRequest("example", nil)
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		c.Do(req)
	// 	}

	// }()
	// go func() {
	// 	req := nando.NewRequest("teste1", nil)

	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		c.Do(req)
	// 	}
	// }()

	// 	go func() {
	// 		c := nando.Client{}
	// 		type Request struct {
	// 			ID          string
	// 			HandlerName string
	// 		}

	// 		req := nando.NewRequest("create-handler", &Request{ID: "1", HandlerName: "NOTIFY"})

	// 		for {
	// 			time.Sleep(5 * time.Second)
	// 			c.Do(req)
	// 		}
	// 	}()

}
