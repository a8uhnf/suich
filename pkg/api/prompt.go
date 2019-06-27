package api

import (
	"github.com/manifoldco/promptui"
)

// RunPrompt generate prompt with given array and label
func RunPrompt(strs []string, label string) (string, error) {
	list := []string{}
	for _, v := range strs {
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