package helps

import "fmt"

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func PanicError(msg string, err error) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func ReturnError(err error) error {
	return err
}
