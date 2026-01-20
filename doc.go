// Package itertools provides lazy sequences similar to Python's itertools,
// implemented using Go 1.23's iter.Seq. All functions accept slices and
// return iter.Seq[T], lazily yielding each result.
//
// Functions included:
//   - Combinations
//   - CombinationsWithReplacement
//   - Permutations
//   - Product
package itertools
