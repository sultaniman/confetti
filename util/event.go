package util

import "fmt"

func PrepareEvent(ref, origin, action string) string {
	return fmt.Sprintf(`ref=%s;origin=%s;action=%s`, ref, origin, action)
}
