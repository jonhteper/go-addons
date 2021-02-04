package addons

// StringAppend permite insertar un string a slice tras comprobar si existe una entrada similar.
func StringAppend(slice []string, addon string) []string {
	var verify bool
	for i := 0; i <= len(slice)-1; i++ {

		if slice[i] == addon {
			verify = true
		}

	}
	if !verify {
		slice = append(slice, addon)
		return slice
	}
	return slice
}

// IntAppend permite insertar un int64 a un slice tras comprobar que es mayor que 0
func IntAppend(slice []int64, addon int64) []int64 {
	if addon > 0 {
		slice = append(slice, addon)
	}
	return slice
}

// FloatAppend permite insertar un float64 a un slice tras comprobar que es mayor que 0
func FloatAppend(slice []float64, addon float64) []float64 {
	if addon > 0 {
		slice = append(slice, addon)
	}
	return slice
}
