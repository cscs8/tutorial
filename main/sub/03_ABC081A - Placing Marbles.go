package sub

import (
	"fmt"
	"strings"
)

/** "Test1 ... aaa." */
func Test1() {
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println(strings.Count(a, "1"))
}
