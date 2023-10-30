package main
// #cgo CFLAGS: -I/usr/include/python3.10
// #cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu -lpython3.10
// #include <Python.h>

import "C"

import (
	"fmt"
    "unsafe"
)

func main(){
    // to initialize interpreter 
    C.Py_Initialize()

    //define module name in c bindings
    moduleName := C.CString("test_wrap")
    //handle memory when function call closes
    defer C.free(unsafe.Pointer(moduleName))
    //import module
    module := C.PyImport_ImportModule(moduleName)
    //error handle
    if module == nil{
        handleError()
        return
    }

    //define function name in C bindings
    functionName := C.CString("add")
    defer C.free(unsafe.Pointer(functionName))
    //import function
    addFunc := C.PyObject_GetAttrString(module, functionName)
    if addFunc == nil{
        handleError()
        return
    }
    //create arguments in C bindings
    arg1 := C.PyLong_FromLong(3)
    arg2 := C.PyLong_FromLong(5)
    defer C.Py_DecRef(arg1)
    defer C.Py_DecRef(arg2)
    args := C.PyTuple_Pack(2, arg1, arg2)
    defer C.Py_DecRef(args)

    //create an object to store the result and call the function
    result := C.PyObject_CallObject(addFunc, args)
    defer C.Py_DecRef(result)
    sum := int(C.PyLong_AsLong(result))
    fmt.Printf("Result: %d\n", sum)

    
    //close interpreter
    C.Py_Finalize()

}

