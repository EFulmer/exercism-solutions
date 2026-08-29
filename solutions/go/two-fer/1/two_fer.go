package twofer

import "fmt"

func ShareWith(name string) string {
    if len(name) == 0 {
        name = "you"
    }
    answer := fmt.Sprintf("One for %s, one for me.", name)
	return answer
}
