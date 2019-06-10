package utils

import "regexp"

func FormatPrice(price *string) {
	r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)

	newPrices := r.FindStringSubmatch(*price)

	if len(newPrices) > 1 {
		*price = newPrices[1]
	} else {
		*price = "Unknown"
	}

}

func FormatStars(stars *string) {
	if len(*stars) >= 3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
}
