package generator

var ReservedWords = map[string]struct{}{
	"Object": {},
	"Any":    {},
}

func PrefixReservedWords(className string) string {
	if _, ok := ReservedWords[className]; ok {
		return "X" + className
	}
	return className
}
