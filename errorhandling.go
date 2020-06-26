package error

import "errors"

func DataNotFound() error {
return errors.New("Data Not Found")
}

func InvalidInput() error {
	return errors.New("Input is invalid")
}

func NoResult() error {
	return errors.New("No result found")
}

func ServerDown() error {
	return errors.New("Server is temporarilr down")
}

func InvalidRequest() error {
	return errors.New("Request is invalid")
}

