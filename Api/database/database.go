package database

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateDB(name string) (err error) {

	folder, err := GetPath()

	if err != nil {
		return err
	}

	filename := fmt.Sprint(folder, name, ".rk")

	_, err = os.Stat(filename)

	if err != nil {
		_, err = os.Create(filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func Write(db string, element map[string]string) (err error) {

	folder, err := GetPath()

	if err != nil {
		return err
	}

	filePath := fmt.Sprint(folder, db, ".rk")

	_, err = os.Stat(filePath)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	var elements []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		lines := scanner.Text()
		elements = strings.Split(lines, "\n")

	}

	for key, val := range element {
		element := fmt.Sprint(key, ":", val)
		elements = append(elements, element)
	}

	_, err = f.WriteString("\n\n" + strings.Join(elements, "\n"))
	if err != nil {
		return err
	}

	return nil

}
