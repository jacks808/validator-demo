package gen

import (
	python3 "github.com/go-python/cpy3"
	"testing"
)

func TestGenerate(t *testing.T) {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	python3.PyRun_SimpleString("print('hello world' )")
}
