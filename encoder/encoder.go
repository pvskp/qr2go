package encoder

import "fmt"

func Encode(data string) string {
	binary_data := ""
	for _, v := range data {
		binary_data += fmt.Sprintf("%08b", v)
	}

	return binary_data
}
