package utility

/*Stolen https://stackoverflow.com/questions/13422578/in-go-how-to-get-a-slice-of-values-from-a-map */
func MapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
