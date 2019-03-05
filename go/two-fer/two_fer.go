// Constructs a simple phrase
package twofer

import "fmt"

// Returns 'One for {name}, one for me.' assings "you" to the name argument
// if none is provided.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
