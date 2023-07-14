package utils

// Deprecated: use maps.Clone() instead
func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

// Deprecated: use maps.Copy like:
// maps.Copy(res, m1)
// maps.Copy(res, m2)
func MergeMap[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	res := make(map[K]V, len(m1)+len(m2))
	for k, v := range m1 {
		res[k] = v
	}

	for k, v := range m2 {
		res[k] = v
	}
	return res
}

// Deprecated: maps.Keys()
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	a := make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

// Deprecated: use maps.Values() instead
func GetMapValues[K comparable, V any](m map[K]V) []V {
	a := make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func GetKeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	kl := make([]K, 0, len(m))
	vl := make([]V, 0, len(m))
	for k, v := range m {
		kl = append(kl, k)
		vl = append(vl, v)
	}
	return kl, vl
}
