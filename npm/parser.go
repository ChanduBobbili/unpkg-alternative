package npm

import "strings"

func ParsePackage(pkg string) (string, string) {
	parts := strings.Split(pkg, "@")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return pkg, "latest"
}