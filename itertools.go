package itertools

import "iter"

func Combinations[T any](items []T, r int) iter.Seq[[]T] { return combinations(items, r) }

func CombinationsWithReplacement[T any](items []T, r int) iter.Seq[[]T] {
	return combinationsWithReplacement(items, r)
}

func Permutations[T any](items []T, r int) iter.Seq[[]T] { return permutations(items, r) }

func Product[T any](items []T, repeat int) iter.Seq[[]T] { return product(items, repeat) }
