package utils

import "github.com/manifoldco/promptui"

// RunPrompt prompts the use with options on the terminal
func RunPrompt(ctxs []string, label string) (string, error) {
	list := []string{}
	for _, v := range ctxs {
		list = append(list, v)
	}
	prompt := promptui.Select{
		Label: label,
		Items: list,
	}
	_, res, err := prompt.Run()
	if err != nil {
		return "", nil
	}
	return res, nil
}
