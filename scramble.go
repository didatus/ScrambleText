package scramble

import (
    "strings"
)

func Scramble(inputString string) string {
    stringElements := strings.Split(inputString, "")
    runes := []rune(inputString)
    outputString := make([]string, len(stringElements))

    for key := range outputString {
        position := int(runes[key]) % len(stringElements)
        outputString[key] = stringElements[position]
        stringElements = append(stringElements[:position], stringElements[position+1:]...)
    }

    return strings.Join(outputString, "")
}
