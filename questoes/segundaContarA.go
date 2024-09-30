package questoes

import (
	"strconv"
	"strings"
)

func CountA(str string) string {
	str = strings.ToLower(str)
	a := strings.Count(str, "a")
	if a == 0 {
		return "a letra A nao aparece nenhuma vez"
	} else {
		return "a letra A aparece " + strconv.Itoa(a) + " vezes"
	}
}
