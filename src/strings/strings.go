package strings

import (
	"fmt"
)

const hi string = "Hi"
const iam string = "iam"
const pl string = "paulo"

func formatString() string {
	return fmt.Sprintf("%s, %s %s", hi, iam, pl)
}

func main() {
	fmt.Printf(formatString())
}
