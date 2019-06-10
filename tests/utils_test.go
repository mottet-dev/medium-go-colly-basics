package tests

import (
	"testing"

	"github.com/mottet-dev/medium-go-colly-basics/utils"
)

func TestFormatPriceWithOnePrice(t *testing.T) {
	input := "$299.00"

	expectedOutput := "299.00"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatPriceWithTwoPrices(t *testing.T) {
	input := "$299.00$480.00"

	expectedOutput := "299.00"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatPriceWithThreePrices(t *testing.T) {
	input := "$299.00$480.00$489.25"

	expectedOutput := "299.00"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatPriceWithFourDigitLenght(t *testing.T) {
	input := "$1234.58"

	expectedOutput := "1234.58"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatPriceWithFourDigitLenghtTwoPrices(t *testing.T) {
	input := "$1234.58$4895.49"

	expectedOutput := "1234.58"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatPriceWithoutCent(t *testing.T) {
	input := "$1234"

	expectedOutput := "1234"

	utils.FormatPrice(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatStars(t *testing.T) {
	input := "4.7 out of 5 stars"

	expectedOutput := "4.7"

	utils.FormatStars(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatStarsTwo(t *testing.T) {
	input := "3.2 out of 5 stars"

	expectedOutput := "3.2"

	utils.FormatStars(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}

func TestFormatStarsThree(t *testing.T) {
	input := "1.0 out of 5 stars"

	expectedOutput := "1.0"

	utils.FormatStars(&input)

	if input != expectedOutput {
		t.Errorf("The input: %s is not equal to the expected output: %s", input, expectedOutput)
	}
}
