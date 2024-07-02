package suffixitsms

import (
	"strings"
)

func (c *client) SendMessage(recipients []string, msg string) (string, error) {
	contacts := prepareRecipientsArray(recipients)
	payload := map[string]string{
		"apiKey":    c.apiKey,
		"text":      msg,
		"toNumbers": contacts,
	}

	body, err := c.sendRequest("POST", "/sendSms", payload)
	if err != nil {
		return "", err
	}

	return body, nil
}

func prepareRecipientsArray(recipients []string) string {
	var numbers []string
	for _, number := range recipients {
		numbers = append(numbers, "88"+cleanPhoneNumber(number))
	}
	return strings.Join(numbers, "+")
}

func cleanPhoneNumber(number string) string {
	number = strings.TrimPrefix(number, "+")
	number = strings.TrimPrefix(number, "88")
	replacer := strings.NewReplacer(" ", "", "-", "", "_", "")
	number = replacer.Replace(number)
	return number
}
