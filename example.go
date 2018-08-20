package main

import (
  "fmt"
  "linip"
)

func main() {
  ini := inifile{path: "example.ini"}
  ini.parse()
  fmt.Println(ini.getvalue("Settings", "StatusPort"))
  fmt.Println(ini.getvalue("FTP", "FTPPort"))
  fmt.Println(ini.getcontainer("FTP")["FTPPort"]) // This is same as code above but twice slower.
  fmt.Println(ini.getmap())
}
