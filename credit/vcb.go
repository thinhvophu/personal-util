package credit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func convertDate(dateStr string) string {
	dateObj, _ := time.Parse("02-01-2006", dateStr)
	return dateObj.Format("02/01")
}

func ConvertVCB(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var outputText []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("converting line: " + line)
		date := convertDate(line[:10])
		line = line[22:]
		descStart := 0
		if line[1] == '-' {
			line = line[2:]
			descStart += 2
		}
		components := strings.Split(line, " - ")
		var amount string
		if strings.Index(components[0], " ") == -1 {
			amount = strings.ReplaceAll(components[0], ",00", "")
		} else {
			amountComponents := strings.Split(components[0], " ")
			amount = amountComponents[len(amountComponents)-1]
		}

		amount = strings.ReplaceAll(amount[:len(amount)-3], ".", ",")
		desc := line[descStart+len(components[0]):]
		desc = strings.ReplaceAll(desc, " - ", "")

		outputText = append(outputText, fmt.Sprintf("%s;%s;%s", date, desc, amount))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(outputText, "\n"), nil
}
