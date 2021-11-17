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
