package main

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// in Go, by default only upper-case (PascalCase) properties can be accessed
// from "outside" the scope of a struct or function, so in this case,
// since JSONs are camelCased because of the JavaScript convention, we need
// to inform (right) what is the field in the JSON body that is going to be
// deserialized

// the functions called inside this one do not return anything, just print to STDOUT what returned from the jsonplaceholder route
func main() {
	// calls get http method - TESTED
	// Get()

	// calls post http method - TESTED
	// Post()

	// calls put http method - TESTED
	Put()

	// calls delete http method - TESTED
	// Delete()
}



