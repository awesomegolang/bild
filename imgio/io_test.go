package imgio

import (
	"bytes"
	"image"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		format  string
		encoder Encoder
		value   image.Image
	}{
		{
			format:  "png",
			encoder: PNGEncoder(),
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			format:  "jpg,jpeg",
			encoder: JPEGEncoder(95),
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		buf := bytes.Buffer{}
		c.encoder(&buf, c.value)
		_, outFormat, err := image.Decode(&buf)
		if err != nil {
			t.Error(err)
		}
		if !strings.Contains(c.format, outFormat) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Encoder", c.format, outFormat)
		}
	}
}
