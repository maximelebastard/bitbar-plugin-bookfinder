package main

import(
  "regexp"
  "log"
)

func replaceAll(src string, regexpString string, replace string) string {
    reg, err := regexp.Compile(regexpString)
    if err != nil {
     log.Fatal(err)
    }

    safe := reg.ReplaceAllString(src, replace)
    return safe
}