package nando

import (
	"sync"
)

var serving = make(map[string]*Server)

type HandleFunc func(*Request) (*Response, error)
type RequestHandler chan *Request
type ReponseHandler chan *Response
type ServerClosed chan bool
type ServerRunning chan bool

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

func NewServer(label string) *Server {
	s := &Server{
		handlers:  make(map[string]HandleFunc),
		requests:  make(RequestHandler),
		responses: make(ReponseHandler),
		closed:    make(ServerClosed),
	}

	serving[label] = s

	return s
}

func (s *Server) Listen(handlers ...*Handler) bool {
	for _, h := range handlers {
		s.HandleFunc(h.Name, h.HandleFunc)
	}
	go s.worker()
	return <-s.closed
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

}
