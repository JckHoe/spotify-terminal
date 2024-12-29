package view

import (
	"fmt"
	"spotify-terminal/internal/view/model"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Caching constant conponents
var emptyRow string
var footerRow string

func renderTitle(builder *strings.Builder, title string, maxLength uint) error {
	titleLen := getStringWidth(title)
	reservedLength := 4
	remainingSpace := maxLength - uint(reservedLength)
	var paddingLeft uint
	var paddingRight uint

	// Validation
	if titleLen >= int(remainingSpace) {
		return fmt.Errorf("Title length of %d is too long for the maxlength", titleLen)
	}
	if builder == nil {
		return fmt.Errorf("Builder is nil")
	}

	// Calculate padding
	remainingLength := remainingSpace - uint(titleLen)
	paddingLeft = remainingLength / 2
	paddingRight = paddingLeft
	if remainingLength%2 != 0 {
		paddingLeft++
	}

	// render title
	builder.WriteString(model.Yellow)
	builder.WriteString(model.CurveTopLeft)
	err := padBuilderWith(builder, paddingLeft, model.Horizontal)
	if err != nil {
		return err
	}
	builder.WriteString(model.SharpLeftDown)
	builder.WriteString(model.Red)
	builder.WriteString(title)
	builder.WriteString(model.Yellow)
	builder.WriteString(model.SharpRightDown)
	err = padBuilderWith(builder, paddingRight, model.Horizontal)
	if err != nil {
		return err
	}
	builder.WriteString(model.CurveTopRight)
	builder.WriteByte('\n')

	return nil
}

func renderItem(builder *strings.Builder, rowContent string, maxLength uint, selected bool) error {
	contentLen := getStringWidth(rowContent)
	reservedLength := 3
	remainingSpace := maxLength - uint(reservedLength)

	// Validation
	if contentLen >= int(remainingSpace) {
		return fmt.Errorf("Content length of %d is too long for the maxlength", contentLen)
	}

	paddingRight := remainingSpace - uint(contentLen)

	// Render item row
	builder.WriteString(model.Yellow)
	builder.WriteString(model.Verticle)
	builder.WriteByte(' ')
	if selected {
		builder.WriteString(model.Green)
	}
	builder.WriteString(rowContent)
	err := padBuilderWith(builder, paddingRight, " ")
	if err != nil {
		return err
	}
	builder.WriteString(model.Yellow)
	builder.WriteString(model.Verticle)
	builder.WriteByte('\n')

	return nil
}

func renderEmptyRow(builder *strings.Builder, maxLength uint) error {
	if emptyRow != "" {
		builder.WriteString(emptyRow)
		return nil
	}

	reservedLength := 2
	remainingSpace := maxLength - uint(reservedLength)

	builder.WriteString(model.Yellow)
	builder.WriteString(model.Verticle)
	err := padBuilderWith(builder, remainingSpace, " ")
	if err != nil {
		return err
	}
	builder.WriteString(model.Verticle)
	builder.WriteByte('\n')

	return nil
}

func renderFooter(builder *strings.Builder, maxLength uint) error {
	if footerRow != "" {
		builder.WriteString(footerRow)
		return nil
	}

	reservedLength := 2
	remainingSpace := maxLength - uint(reservedLength)

	builder.WriteString(model.Yellow)
	builder.WriteString(model.CurveBtmLeft)
	err := padBuilderWith(builder, remainingSpace, model.Horizontal)
	if err != nil {
		return err
	}
	builder.WriteString(model.CurveBtmRight)

	return nil
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

func getStringWidth(s string) int {
	width := 0
	for _, r := range s {
		width += runewidth.StringWidth(string(r))
	}
	return width
}
