package main

import "fmt"

const (
	french = "french"
	tamazight = "tamazight"

	arabicPrefix = "Salam, "
	tamazightPrefix = "Azul, "
	frenchPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "3alam"
	}

	return gettingPrefix(language) + name + "!"
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case tamazight:
		prefix = tamazightPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = arabicPrefix
	} 
	return
}

func main() {
	fmt.Println(Hello("hayat", ""))
}
