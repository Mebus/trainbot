// Package testutil contains some utils for testing images.
package testutil

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func iabs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// AssertImagesAlmostEqual asserts that two images are almost equal.
func AssertImagesAlmostEqual(t *testing.T, truth image.Image, test image.Image) {
	require.Equal(t, truth.Bounds().Size(), test.Bounds().Size())

	var diff uint64
	for y := 0; y < truth.Bounds().Dy(); y++ {
		for x := 0; x < truth.Bounds().Dx(); x++ {
			r0, g0, b0, _ := truth.At(x, y).RGBA()
			r1, g1, b1, _ := test.At(x, y).RGBA()

			diff += uint64(iabs(int(r0)-int(r1)) / 255)
			diff += uint64(iabs(int(g0)-int(g1)) / 255)
			diff += uint64(iabs(int(b0)-int(b1)) / 255)
		}
	}

	diffPerPx := float64(diff) / float64(truth.Bounds().Dx()) / float64(truth.Bounds().Dy()) / 3
	assert.Less(t, diffPerPx, 1.)
}
