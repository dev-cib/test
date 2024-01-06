import "fmt"

// this is also "package block" scope
// this is exported because the identified "Fascinating" is capitalized
func Fascinating() {
	fmt.Println("Planet speed:", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet spinning speed", planetRotationSpeed)
}