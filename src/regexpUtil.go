package main

import(
  "regexp"
)

func replaceAll(src string, regexpString string, replace string) string {
    reg, err := regexp.Compile(regexpString)
    deal(err)

    safe := reg.ReplaceAllString(src, replace)
    return safe
}