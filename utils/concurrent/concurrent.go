package concurrent

// WaitForError wait chan error, returns first error.
func WaitForError(ch chan error, chLen int) error {
	var err error
	var ret error
	for i := 0; i < chLen; i++ {
		ret = <-ch
		if err == nil {
			err = ret
		}
	}
	return err
}