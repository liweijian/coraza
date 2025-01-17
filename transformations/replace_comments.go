// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package transformations

func replaceComments(data string) (string, error) {
	return doReplaceComments(data), nil
}

func doReplaceComments(value string) string {
	var i, j int
	incomment := false

	input := []byte(value)
	inputLen := len(input)
	for i < inputLen {
		if !incomment {
			if (input[i] == '/') && (i+1 < inputLen) && (input[i+1] == '*') {
				incomment = true
				i += 2
			} else {
				input[j] = input[i]
				i++
				j++
			}
		} else {
			if (input[i] == '*') && (i+1 < inputLen) && (input[i+1] == '/') {
				incomment = false
				i += 2
				input[j] = ' '
				j++
			} else {
				i++
			}
		}
	}

	if incomment {
		input[j] = ' '
		j++
	}

	return string(input[0:j])
}
