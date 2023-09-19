package nando

import "sync"

var serving = NewServer()

type HandleFunc func(*Request) (*Response, error)
type RequestHandler chan *Request
type ReponseHandler chan *Response
type ServerClosed chan bool

type Handler struct {
	Name       string
	HandleFunc HandleFunc
}

func NewHandler(name string, h HandleFunc) *Handler { return &Handler{Name: name, HandleFunc: h} }

type Server struct {
	requests  RequestHandler
	responses ReponseHandler
	handlers  map[string]HandleFunc
	sync.Mutex
	closed ServerClosed
}

func NewServer() *Server {
	return &Server{
		handlers:  make(map[string]HandleFunc),
		requests:  make(RequestHandler),
		responses: make(ReponseHandler),
		closed:    make(ServerClosed),
	}
}

func Listen(handlers ...*Handler) bool {
	for _, h := range handlers {
		serving.HandleFunc(h.Name, h.HandleFunc)
	}
	go serving.worker()
	return <-serving.closed
}

func (s *Server) HandleFunc(funcName string, handler HandleFunc) {
	s.handlers[funcName] = handler
}

func (s *Server) Submit(req *Request) {
	s.Lock()
	defer s.Unlock()
	select {
	case <-s.closed:
		return
	default:
		s.requests <- req

	}
}

func (s *Server) Close() {
	s.Lock()
	defer s.Unlock()

	s.closed <- true

	close(s.requests)
}

func (s *Server) worker() {

	for {
		req := <-s.requests
		resp, err := s.handlers[req.funcName](req)
		if err != nil {
			panic(err)
		}

		s.responses <- resp
	}
	// for req := range s.requests {
	// 	resp, err := s.handlers[req.funcName](req)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	s.responses <- resp
	// }
}
