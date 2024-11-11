package utils

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BoolToEnumString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
func EnumStringToBool(value string) bool {
	switch value {
	case "true", "on", "yes", "1":
		return true
	case "false", "off", "no", "0":
		return false
	default:
		// Retorna false por padrão ou pode levantar um erro dependendo do caso
		return false
	}
}
func EnumStringPointerToBool(value *string) bool {
	if value == nil {
		// Se o ponteiro é nulo, retorna false por padrão ou pode levantar um erro dependendo do caso
		return false
	}

	switch *value {
	case "true", "on", "yes", "1":
		return true
	case "false", "off", "no", "0":
		return false
	default:
		// Retorna false por padrão ou pode levantar um erro dependendo do caso
		return false
	}
}
func GetStringOrDefault(ptr *string, defaultVal string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// GetIntOrDefault retorna o valor do ponteiro se não for nil, caso contrário retorna um valor padrão
func GetIntOrDefault(ptr *int, defaultVal int) int {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}
