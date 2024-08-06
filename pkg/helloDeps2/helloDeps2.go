package helloDeps2

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps2.PrintPhrase")
	return phrase
}
