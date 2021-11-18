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

func Read(db string) (result map[string]string, err error) {

	folder, err := GetPath()

	if err != nil {
		return result, err
	}

	filePath := fmt.Sprint(folder, db, ".rk")

	_, err = os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return result, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	result = map[string]string{}

	for scanner.Scan() {

		line := scanner.Text()

		value := strings.Split(line, ":")
		if len(value) >= 2 {
			result[value[0]] = value[1]
		}

	}

	if len(result) <= 0 {
		err = fmt.Errorf("no result found : %v", filePath)

		return result, err
	}

	return result, err
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
		elements = append(elements, line)

	}

	for key, val := range element {
		stringToAppend = fmt.Sprint(key, ":", val)
		elements = append(elements, stringToAppend)
	}

	_, err = f.WriteString("\n" + strings.Join(elements, "\n") + "\n")
	if err != nil {
		return err
	}

	return nil

}

// func Delete(db string, element map[string]string) (err error) {

// 	elements, err := Read(db)

// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("avant : ", elements)

// 	for i, el := range elements {
// 		for k1, v1 := range el {
// 			for k2, v2 := range element {
// 				if k1 == k2 && v1 == v2 {
// 					Cut(elements, i)
// 				}
// 			}
// 		}

// 	}

// 	fmt.Println("après : ", elements)

// 	return err
// }

func DeleteDB(db string) (err error) {

	folder, err := GetPath()

	if err != nil {
		return err
	}

	filePath := fmt.Sprint(folder, db, ".rk")

	_, err = os.Stat(filePath)

	if err != nil {
		return nil
	}

	err = os.Remove(filePath)

	return err
}
