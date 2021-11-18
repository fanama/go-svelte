package database

import "os"

func GetPath() (result string, err error) {

	_, err = os.Stat("data/")

	if err != nil {
		_, err = os.Stat("../data/")
		if err != nil {
			return "../data/", err
		}
		return "../data/", err
	}

	return "data/", err
}

func Cut(a []map[string]string, i int) []map[string]string {

	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = nil  // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.

	return a
}
