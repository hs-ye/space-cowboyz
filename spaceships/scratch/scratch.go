package scratch

import "fmt"

type TestStruct struct {
	Data string
}

func (t TestStruct) Dostuff(text string) (t2 TestStruct) {
	t.Data = "I made it do " + text
	return t
}

func (t *TestStruct) DoPointerstuff(text string) (t2 TestStruct) {
	t.Data = "I made this pointer do " + text // but here no need to pass in address...compiler does the addressing & operation for you?
	return *t                                 // should return the pointer?? or the copy itself?
}

// scratch pad testing
// trying to figure out the difference between a receiver vs a pointer receiver
// is it literally modify in place, vs modify a copy?
func ScratchTest() {
	t1 := TestStruct{Data: "original text"}
	p1 := &t1

	fmt.Printf("The setup: pointer %s vs value %s", p1, t1)

	fmt.Println("Doing stuff on value receiver")
	fmt.Println(t1.Dostuff("stuff").Data) // result of function, is a copy
	fmt.Println(t1.Data)                  // original not modified

	fmt.Println("Doing stuff on Pointer receiver")
	fmt.Println(t1.DoPointerstuff("pointer stuff using values").Data) // pointer modified
	fmt.Println(t1.Data)                                              // original is now modified
	fmt.Println(*p1)                                                  // original is now modified
	fmt.Println(p1.Data)                                              // original is now modified
	fmt.Println("Now for pointer version")
	fmt.Println(p1.DoPointerstuff("pointer stuff using receiver").Data)
	fmt.Println(p1)  // Pointer version
	fmt.Println(*p1) // still pionter version
	fmt.Println("Resetting data")
	t1 = TestStruct{Data: "reset to original text"}
	fmt.Println(*p1)                     // still pionter version
	fmt.Println(p1.Dostuff("test").Data) // passing a pointer to the do function
	fmt.Println(p1.Data)                 // doesn't do anything, cos it was a copy

}
