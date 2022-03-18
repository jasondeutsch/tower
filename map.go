package tower

// Map takes a slice generic type A, a function to move from type A to type B
// where A and B can be different types
// and returns a slice of type B.
//
// Usage:
// Map([]int{1,2,3},)
func Map[A, B any](a []A, fn func(A) B) []B {
	b := make([]B, len(a))
	for i := range a {
		b[i] = fn(a[i])
	}
	return b
}
