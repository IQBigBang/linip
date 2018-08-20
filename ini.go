package linip

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type inifile struct {
	path string
	data map[string]string
}

// if measuretime is true, parsing time will be printed
func (ini *inifile) parse(measuretime bool) {
	if measuretime {
		defer timeTrack(time.Now())
	}
	fileInfo, err := os.Stat(ini.path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		} else {
			log.Fatal("File cannot be opened: ")
			log.Println(fileInfo)
		}
	} else {
		file, err := os.Open(ini.path)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		eof := false
		var currentContainer string
		ini.data = make(map[string]string)
		for !eof {
			success := scanner.Scan()
			if success == false {
				//False: error or EOF
				err = scanner.Err()
				if err == nil {
					eof = true
				} else {
					log.Fatal(err)
				}
			} else {
				//PARSING
				currentLine := scanner.Text()
				currentLine = strings.TrimSpace(currentLine)
				if strings.HasPrefix(currentLine, "[") {
					//Container
					currentContainer = currentLine[1:(len(currentLine) - 1)]
				} else if strings.HasPrefix(currentLine, ";") {
					//Comment
				} else if currentLine == "" {
					//Empty line
				} else {
					index := strings.IndexByte(currentLine, '=')
					varname := currentLine[:index]
					varname = strings.TrimSpace(varname)
					varvalue := currentLine[(index + 1):]
					commentIndex := strings.IndexByte(varvalue, ';')
					if commentIndex != -1 {
						varvalue = varvalue[:commentIndex] //Strip comment
					}
					varvalue = strings.TrimSpace(varvalue)
					ini.data[currentContainer+"."+varname] = varvalue
				}
			}
		}
	}
}

func (ini *inifile) getvalue(containername string, varname string) string {
	return ini.data[containername+"."+varname]
}

func (ini *inifile) getcontainer(containername string) map[string]string {
	returnMap := make(map[string]string)
	for key := range ini.data {
		if strings.HasPrefix(key, containername+".") {
			returnMap[key] = ini.data[key]
		}
	}
	return returnMap
}

func (ini *inifile) getmap() map[string]string {
	return ini.data
}

//From Coderwall
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Parsing took %s", elapsed)
}
