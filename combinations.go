package itertools

import "iter"

func combinations[T any](items []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(items)
		if r < 0 {
			r = n
		}

		if r == 0 {
			yield([]T{})
			return
		}
		if r > n {
			return
		}

		inds := make([]int, r)
		for i := range inds {
			inds[i] = i
		}

		for {
			out := make([]T, r)
			for i := range inds {
				out[i] = items[inds[i]]
			}
			if !yield(out) {
				return
			}

			i := r - 1
			for i >= 0 && inds[i] == i+n-r {
				i--
			}
			if i < 0 {
				return
			}

			inds[i]++
			for j := i + 1; j < r; j++ {
				inds[j] = inds[j-1] + 1
			}
		}
	}
}

func combinationsWithReplacement[T any](items []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(items)
		if r == 0 {
			yield([]T{})
			return
		}

		inds := make([]int, r)

		for {
			out := make([]T, r)
			for i := range inds {
				out[i] = items[inds[i]]
			}
			if !yield(out) {
				return
			}

			i := r - 1
			for i >= 0 && inds[i] == n-1 {
				i--
			}
			if i < 0 {
				return
			}

			inds[i]++
			for j := i + 1; j < r; j++ {
				inds[j] = inds[j-1]
			}
		}
	}
}
