package main
// #cgo CFLAGS: -I/usr/include/python3.8
// #cgo LDFLAGS: -lpython3.8
// #include <Python.h>

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func initialize() {
    // Initialize Python
    C.Py_Initialize()

    // Import the Python module
    module := C.PyImport_ImportModule("my_python_module")

    // Check for errors when importing
    if module == nil {
        C.PyErr_Print()
        return
    }

    // Get a reference to the Python function
    function := C.PyObject_GetAttrString(module, "my_python_function")

    // Check for errors when getting the function
    if function == nil {
        C.PyErr_Print()
        return
    }

    // Prepare and call the Python function
    args := C.Py_BuildValue("(s)", "Hello, Python!")
    result := C.PyObject_CallObject(function, args)

    // Check for errors when calling the function
    if result == nil {
        C.PyErr_Print()
        return
    }

    // Handle the result (convert it to Go data if necessary)

    // Finalize Python
    C.Py_Finalize()
    fmt.Println("initialized")
}

func main() {
    initialize
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go!")
    })

    handler := cors.Default().Handler(mux)

    http.ListenAndServe(":8080", handler)
}
