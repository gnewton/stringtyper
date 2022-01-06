# stringtyper

stringtype tries to determimine the Go type that will best satisfy
representing a series of one or more example strings supplied by the
user.

Where "best" is the most specific Go type that can represent what is
presented, using the range of the type where appropriate. 
So for int and float types, the range can constrain the type that is
selected, with int and float type ranges defined in
[https://pkg.go.dev/builtin](https://pkg.go.dev/builtin) and [https://pkg.go.dev/math#pkg-constants](https://pkg.go.dev/math#pkg-constants).
Because of this, the example strings presented need to represent both
the type and the range that is to be used. 

## Float formats
Internally,
[strconv.ParseFloat](https://pkg.go.dev/strconv#ParseFloat) is used,
and the following float fmt
[formats](https://pkg.go.dev/fmt#hdr-Printing) are supported: 

`%e %E %f %g %G %x %X`

and the following fmt format is not supported:

`%b`


For examples, see the tests in
[https://github.com/gnewton/stringtyper/blob/main/pkg/stringtyper/stringtyper_test.go](https://github.com/gnewton/stringtyper/blob/main/pkg/stringtyper/stringtyper_test.go).

# Usage
From [examples/simple/main.go](https://github.com/gnewton/stringtyper/blob/main/pkg/stringtyper/examples/simple/main.go)

    package main
    
    import (
    	"fmt"
    	"github.com/gnewton/stringtyper/pkg/stringtyper"
    )
    
    func main() {
    	var examples []string
    
    	// bool
    	examples = []string{"true", "true", "false"}
    	ti := stringtyper.NewStringTyper()
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// string
    	examples = []string{"true", "true", "27"}
    	ti = stringtyper.NewStringTyper()
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// int8
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"-127", "120", "27"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// uint8
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"9", "127", "27"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// uint8
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"9", "129", "27"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// uint16
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"9", "129", "27", "256"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// float32
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"-9", "129.5", "27", "256"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// float64
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"-9", "129.5", "27", "256", "3.40282346638528859811704183484516925440e+39"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    
    	// string
    	ti = stringtyper.NewStringTyper()
    	examples = []string{"u", "-9", "129.5", "27", "256", "3.40282346638528859811704183484516925440e+39"}
    	for _, ex := range examples {
    		ti.CheckFieldTypeAndLength(ex)
    	}
    	fmt.Printf("Type is: %s %s\n", ti.Kind(), examples)
    }
