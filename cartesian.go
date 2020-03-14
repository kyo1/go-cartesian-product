package cartesian

import (
	"context"
)

func All(ctx context.Context, set []interface{}) chan []interface{} {
	ch := make(chan []interface{})

	go func() {
		defer close(ch)

		pos := make([]int, 1)

		for {
			if ctx != nil {
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			pair := make([]interface{}, len(pos))
			for i, p := range pos {
				pair[len(pos)-i-1] = set[p] // dictionary order
			}
			ch <- pair

			for i := 0; i < len(pos); i++ {
				pos[i]++ // pos[i] = (pos[i] + 1) % len(set) is slow
				if pos[i] != len(set) {
					break
				}
				pos[i] = 0
				if i == len(pos)-1 {
					pos = append(pos, 0)
					break
				}
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
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			pair := make([]interface{}, repeat)
			for i, p := range pos {
				pair[repeat-i-1] = set[p] // dictionary order
			}
			ch <- pair

			cnt := 0
			for i := 0; i < repeat; i++ {
				pos[i]++ // pos[i] = (pos[i] + 1) % len(set) is slow
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
