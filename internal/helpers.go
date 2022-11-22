package internal

import "fmt"

func ResolveAddress(url string, resource string) string {
	return fmt.Sprintf("%s/%s", url, resource)
}
