package pmatch

import (
	"image"
)

func SearchGrayOpt(img, pat *image.Gray) (maxX, maxY int, maxScore float64) {
	if pat.Bounds().Size().X > img.Bounds().Size().X ||
		pat.Bounds().Size().Y > img.Bounds().Size().Y {
		panic("patch too large")
	}

	// search rect in img coordinates
	searchRect := image.Rectangle{
		Min: img.Bounds().Min,
		Max: img.Bounds().Max.Sub(pat.Rect.Size()),
	}

	m, n := searchRect.Dx(), searchRect.Dy()
	du, dv := pat.Rect.Dx(), pat.Rect.Dy()

	is := img.Stride
	ps := pat.Stride

	imgX0, imgY0 := img.Rect.Min.X, img.Rect.Min.Y
	patX0, patY0 := pat.Rect.Min.X, pat.Rect.Min.Y

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			winX0 := imgX0 + x
			winY0 := imgY0 + y
			imgPatStartIx := y*is + x

			var dot, sqSumI, sqSumP uint64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					pxIi := ((winY0+v)-winY0)*is + ((winX0+u)-winX0)*1
					pxI := img.Pix[imgPatStartIx+pxIi]

					pxPi := ((patY0+v)-patY0)*ps + ((patX0+u)-patX0)*1
					pxP := pat.Pix[pxPi]

					dot += uint64(pxI) * uint64(pxP)
					sqSumI += uint64(pxI) * uint64(pxI)
					sqSumP += uint64(pxP) * uint64(pxP)
				}
			}

			abs := float64(sqSumI) * float64(sqSumP)
			var score float64
			if abs == 0 {
				score = 1
			} else {
				score = float64(dot*dot) / abs
			}

			if score > maxScore {
				maxScore = score
				maxX, maxY = x, y
			}
		}
	}

	return
}