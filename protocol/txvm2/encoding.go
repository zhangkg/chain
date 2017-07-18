package txvm2

import (
	"encoding/binary"
	"fmt"
)

func opEncode(vm *vm) {
	s := vm.popBytes(datastack)
	vm.push(datastack, vbytes(encode(s)))
}

func opInt64(vm *vm) {
	a := vm.popBytes(datastack)
	res, n := binary.Varint(a)
	if n <= 0 {
		panic(fmt.Errorf("int64: not a valid varint: %x", a))
	}
	vm.push(datastack, vint64(res))
}

func smallInt(n vint64) func(*vm) {
	return func(vm *vm) {
		vm.push(datastack, n)
	}
}

func opPushdata(vm *vm) {
	// xxx
}