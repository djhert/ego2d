package ego

import "fmt"

const (
	NAME       = "eGo Game Engine 2D"
	MAJORVER   = 0
	MINORVER   = 1
	RELEASEVER = 1
)

func Version() string {
	return fmt.Sprintf("%s - v%d.%d.%d", NAME, MAJORVER, MINORVER, RELEASEVER)
}
