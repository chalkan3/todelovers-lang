package engine

import "io/ioutil"

type file struct{}

func (f *file) open(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	content := string(data)
	return content, nil
}
func newFile() *file { return &file{} }

func File(filename string) (string, error) { return newFile().open(filename) }
