package itertools

import "iter"

func product[T any](items []T, repeat int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if repeat <= 0 {
			yield([]T{})
			return
		}

		n := len(items)
		if n == 0 {
			return
		}

		inds := make([]int, repeat)

		for {
			out := make([]T, repeat)
			for i := range inds {
				out[i] = items[inds[i]]
			}
			if !yield(out) {
				return
			}

			i := repeat - 1
			for i >= 0 && inds[i] == n-1 {
				i--
			}
			if i < 0 {
				return
			}

			inds[i]++
			for j := i + 1; j < repeat; j++ {
				inds[j] = 0
			}
		}
	}
}
