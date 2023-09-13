package memory

type Page struct {
	data     []byte
	frame    int
	dirty    bool
	accessed bool
	next     *Page
}
