package gojsonld

import (
	"io/ioutil"
)

/*
	Reads an RDF file from disk and converts it into a dataset object
	args:
		path: the path to the RDF file
	returns:
		a tuple consisting of the dataset and an error code. If the execution was successful,
		the error code is nil
*/
func ReadDatasetFromFile(path string) (*Dataset, error) {
	file, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		return nil, fileErr
	}
	dataset, parseErr := ParseDataset(file)
	return dataset, parseErr
}
