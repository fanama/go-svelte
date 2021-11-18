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
		if err = CreateDB(db); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	var elements []string

	var stringToAppend string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		fmt.Println("line :", line)
		elements = append(elements, line)

	}

	for key, val := range element {
		stringToAppend = fmt.Sprint(key, ":", val)
		elements = append(elements, stringToAppend)
	}

	for i, e := range elements {
		fmt.Println(i, e)
	}

	_, err = f.WriteString("\n" + strings.Join(elements, "\n"))
	if err != nil {
		return err
	}

	return nil

}
