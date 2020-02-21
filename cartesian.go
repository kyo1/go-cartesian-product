package cartesian

import (
	"context"
)

func All(ctx context.Context, set []interface{}) chan []interface{} {
	ch := make(chan []interface{})

	go func() {
		defer close(ch)

		sz := 1
		pos := make([]int, sz)

		for {
			if ctx != nil {
				// stop generator, if context was cancelled
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			// construct pair from set
			pair := make([]interface{}, sz)
			for i, p := range pos {
				pair[i] = set[p]
			}
			ch <- pair

			for i := 0; i < sz; i++ {
				pos[i]++
				if pos[i] != len(set) {
					break
				}
				if i == sz-1 {
					pos = append(pos, 0)
					sz++
				}
				pos[i] = 0
			}
		}
	}()

	return ch
}

func Product(ctx context.Context, set []interface{}, repeat int) chan []interface{} {
	ch := make(chan []interface{})

	go func() {
		defer close(ch)

		pos := make([]int, repeat)

		for {
			if ctx != nil {
				// stop generator, if context was cancelled
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			// construct pair from set
			pair := make([]interface{}, repeat)
			for i, p := range pos {
				pair[i] = set[p]
			}
			ch <- pair

			cnt := 0
			for i := 0; i < repeat; i++ {
				pos[i]++
				if pos[i] != len(set) {
					break
				}
				pos[i] = 0
				cnt++
			}

			if cnt == repeat {
				break
			}
		}
	}()

	return ch
}
