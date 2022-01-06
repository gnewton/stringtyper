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
