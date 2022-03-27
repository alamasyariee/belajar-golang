package salam

import "fmt"

func Salam(name string) string {
	message := fmt.Sprintf("Assalamu 'alaikum %v. Semoga Allah merahmatimu", name)

	return message
}
