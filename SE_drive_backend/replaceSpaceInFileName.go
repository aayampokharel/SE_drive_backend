package main

import "strings"

func replaceSpaceInFileName(fileName string) string {

	return strings.ReplaceAll(fileName, " ", "_")
}
