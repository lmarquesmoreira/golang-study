package exercises

// MyReader example
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (t MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}
