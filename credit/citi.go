package credit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertCiti(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var outputText []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		components := strings.Fields(line)
		date := components[0]
		desc := strings.Join(components[1:len(components)-1], " ")
		amount := components[len(components)-1]
		if strings.HasPrefix(amount, "(") && strings.HasSuffix(amount, ")") {
			amount = "-" + strings.Trim(amount, "()")
		}
		outputText = append(outputText, fmt.Sprintf("%s;%s;%s", date, desc, amount))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(outputText, "\n"), nil
}
