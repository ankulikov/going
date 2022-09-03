package going

type Stream[T any] <-chan T

func NewStream[T any](items ...T) Stream[T] {
	out := make(chan T)
	go func() {
		for _, item := range items {
			out <- item
		}
		close(out)
	}()
	return out
}

func (s Stream[T]) Filter(predicate func(arg T) bool) Stream[T] {
	out := make(chan T)
	go func() {
		for e := range s {
			if predicate(e) {
				out <- e
			}
		}
		close(out)
	}()
	return out
}

func (s Stream[T]) Map(transformer func(item T) T) Stream[T] {
	out := make(chan T)
	go func() {
		for e := range s {
			out <- transformer(e)
		}
		close(out)
	}()
	return out
}

func (s Stream[T]) Count() int {
	counter := 0
	for range s {
		counter++
	}
	return counter
}

func (s Stream[T]) ToSlice() []T {
	var out = make([]T, 0)
	for e := range s {
		out = append(out, e)
	}
	return out
}
