package typeAst

// Check for skalar types
// https://developers.google.com/protocol-buffers/docs/proto3#scalar
func IsScalar(typename string) bool {
	switch typename {
	case "double":
		return true
	case "float":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "sint64":
		return true
	case "fixed32":
		return true
	case "fixed64":
		return true
	case "sfixed32":
		return true
	case "sfixed64":
		return true
	case "bool":
		return true
	case "string":
		return true
	case "bytes":
		return true
	}
	return false
}
