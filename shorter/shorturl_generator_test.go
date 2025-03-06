package shorter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://github.com/Xavier-Tse/go-tinyurl"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://github.com/Xavier-Tse/Gedis"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://github.com/Xavier-Tse/go-mysql-elasticsearch"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "99aDzbJM")
	assert.Equal(t, shortLink_2, "eTLsYrbr")
	assert.Equal(t, shortLink_3, "AtFehFY1")
}
