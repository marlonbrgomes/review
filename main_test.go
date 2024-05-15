func TestHelloWorld(t *testing.T) {
	// Redirect standard output to a buffer
	buf := &bytes.Buffer{}
	old := os.Stdout
	os.Stdout = buf
	defer func() {
		os.Stdout = old
	}()

	// Call the helloWorld function
	helloWorld()

	// Get the output from the buffer
	output := buf.String()

	// Check if the output matches the expected value
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}