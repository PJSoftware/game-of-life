package life

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type INI struct {
	section  map[string]bool
	value    map[string]map[string]string
	fileName string
}

func (i *INI) Parse(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	i.section = make(map[string]bool)
	i.value = make(map[string]map[string]string)

	i.fileName = fileName
	patternSect := `^[[](.+)[]]`
	patternValue := `^(\S+)\s*=\s*(\S.*)$`
	currSect := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		isSect, err := regexp.MatchString(patternSect, line)
		if err != nil {
			return err
		}
		if isSect {
			reSect, err := regexp.Compile(patternSect)
			if err == nil {
				currSect = reSect.FindStringSubmatch(line)[1]
				i.section[currSect] = true
				i.value[currSect] = make(map[string]string)
			} else {
				log.Panic("INI: Could not compile regexp to read Section")
			}
		} else {
			isValue, err := regexp.MatchString(patternValue, line)
			if err != nil {
				return err
			}
			if isValue {
				reVal, err := regexp.Compile(patternValue)
				if err == nil {
					key := reVal.FindStringSubmatch(line)[1]
					value := reVal.FindStringSubmatch(line)[2]
					i.value[currSect][key] = value
				} else {
					log.Panic("INI: Could not compile regexp to read Value")
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (i *INI) Section(sectName string) error {
	if i.section[sectName] {
		return nil
	} else {
		return fmt.Errorf("INI: Section '%s' not found in '%s'", sectName, i.fileName)
	}
}

func (i *INI) Value(sectName string, valName string) (string, error) {
	err := i.Section(sectName)
	if err != nil {
		return "", err
	}

	if val, ok := i.value[sectName][valName]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("INI: Value '%s' not found in section '%s' of '%s'", valName, sectName, i.fileName)
	}
}
