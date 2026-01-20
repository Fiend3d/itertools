package itertools

import "iter"

func permutations[T any](items []T, r int) iter.Seq[[]T] {
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

		used := make([]bool, n)
		curr := make([]T, 0, r)

		var dfs func() bool
		dfs = func() bool {
			if len(curr) == r {
				out := make([]T, r)
				copy(out, curr)
				return yield(out)
			}

			for i := range n {
				if used[i] {
					continue
				}
				used[i] = true
				curr = append(curr, items[i])
				if !dfs() {
					return false
				}
				curr = curr[:len(curr)-1]
				used[i] = false
			}
			return true
		}

		dfs()
	}
}
