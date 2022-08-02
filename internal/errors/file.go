package hferrors

import "fmt"

type FileExistsError struct {
	Filename string
}

func NewFileExistsError(filename string) *FileExistsError {
	return &FileExistsError{
		Filename: filename,
	}
}

func (err *FileExistsError) Error() string {
	return fmt.Sprintf(
		"error: file %s already exists. Please use a different filename.",
		err.Filename,
	)
}

type FailedFileCreationError struct {
	filename string
	context  error
}

func NewFailedFileCreationError(filename string) *FailedFileCreationError {
	return &FailedFileCreationError{
		filename: filename,
	}
}

func (err *FailedFileCreationError) Error() string {
	return fmt.Sprintf(
		"error: failed to create csv file %s: %s",
		err.filename,
		err.context,
	)
}

func (err *FailedFileCreationError) Wrap(wrapped error) *FailedFileCreationError {
	err.context = wrapped
	return err
}
