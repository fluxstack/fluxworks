package maps

import "github.com/samber/lo"

func Keys[K comparable, V any](v map[K]V) []K {
	return lo.Keys(v)
}
