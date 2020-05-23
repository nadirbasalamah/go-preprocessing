package preprocessing

import "fmt"

func ExamplePreprocessing() {
	xs := []string{"this", "is", "the", "example", "!", "this", "is", "great"}
	fmt.Println(Preprocessing(xs))
	// Output:
	// [  the example  this is great]
}
