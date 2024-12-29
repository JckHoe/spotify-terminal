package view

import (
	"fmt"
	"spotify-terminal/internal/view/model"
	"strings"
)

func renderTitle(title string, maxLength uint) (string, error) {
	// Variables
	titleLen := len(title)
	reservedLength := 4
	remainingSpace := maxLength - uint(reservedLength)
	var titleBuilder strings.Builder
	var paddingLeft uint
	var paddingRight uint

	// Validation
	if titleLen >= int(remainingSpace) {
		return "", fmt.Errorf("Title length of %d is too long for the maxlength", titleLen)
	}

	// Calculate padding
	remainingLength := remainingSpace - uint(titleLen)
	paddingLeft = remainingLength / 2
	paddingRight = paddingLeft
	if remainingLength%2 != 0 {
		paddingLeft++
	}

	// render title
	titleBuilder.WriteString(model.Yellow)
	titleBuilder.WriteString(model.CurveTopLeft)
	err := padBuilderWith(&titleBuilder, paddingLeft, model.Horizontal)
	if err != nil {
		return "", err
	}
	titleBuilder.WriteString(model.SharpLeftDown)
	titleBuilder.WriteString(title)
	titleBuilder.WriteString(model.Yellow)
	titleBuilder.WriteString(model.SharpRightDown)
	err = padBuilderWith(&titleBuilder, paddingRight, model.Horizontal)
	if err != nil {
		return "", err
	}
	titleBuilder.WriteString(model.CurveTopRight)
	titleBuilder.WriteByte('\n')
	titleBuilder.WriteByte('\n')

	return titleBuilder.String(), nil
}

func padBuilderWith(builder *strings.Builder, length uint, padding string) error {
	if padding == "" {
		return fmt.Errorf("Empty padding characters passed")
	}

	if builder == nil {
		return fmt.Errorf("String Builder is nil")
	}

	for i := 0; i < int(length); i++ {
		builder.WriteString(padding)
	}

	return nil
}
