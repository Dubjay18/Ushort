package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	g := NewGenerator()
	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1 := g.GenerateShortLink(initialLink_1)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := g.GenerateShortLink(initialLink_2)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := g.GenerateShortLink(initialLink_3)

	assert.Equal(t, shortLink_1, "X96ZgU63")
	assert.Equal(t, shortLink_2, "U3ecosYc")
	assert.Equal(t, shortLink_3, "YKLN34Ab")
}
