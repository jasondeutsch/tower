package tower

func Filter[A any](fn func(A) bool, a []A) []A {
	var b []A
	for i := range a {
		if fn(a[i]) {
			b = append(b, a[i])
		}
	}
	return b
}
