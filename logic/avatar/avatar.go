package avatar

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"apihut-server/config"
	"apihut-server/logic/consts"
	"apihut-server/models"

	"github.com/nullrocks/identicon"
)

func NewAvatar(req *models.AvatarReq) (string, error) {
	gen, err := identicon.New(
		req.GetNamespace(),
		req.GetBlock(),
		req.GetDensity(),
		identicon.SetBackgroundColorFunction(func(cb []byte, fc color.Color) color.Color {
			if len(req.BackgroundColor) != 0 {
				// 透明
				if req.BackgroundColor == consts.ColorTransparent.String() {
					return color.Transparent
				}
				// 自定义背景颜色
				backgroundColor, err := hexToRGBA(req.BackgroundColor)
				if err != nil {
					return defaultBackgroundColor()
				}
				return backgroundColor
			}
			// 默认颜色
			return defaultBackgroundColor()
		}),
		identicon.SetFillColorFunction(func(hashBytes []byte) color.Color {
			// 自定义颜色
			if len(req.FillColor) != 0 {
				fillColor, err := hexToRGBA(req.FillColor)
				if err != nil {
					return defaultFillColor(hashBytes)
				}
				return fillColor
			}
			// 默认颜色
			return defaultFillColor(hashBytes)
		}),
	)
	if err != nil {
		return "", err
	}
	ident, err := gen.Draw(req.GetHash())
	if err != nil {
		return "", err
	}

	filePath := getFilePath(req)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return "", err
	}

	switch req.GetOutput() {
	case consts.JPG, consts.JPEG:
		err = ident.Jpeg(req.GetSize(), req.GetQuality(), f)
	case consts.SVG:
		err = ident.Svg(req.GetSize(), f)
	default:
		err = ident.Png(req.GetSize(), f)
	}
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func getFilePath(req *models.AvatarReq) string {
	ext := req.GetOutput().String()
	switch req.GetOutput() {
	case consts.JPG, consts.JPEG:
		ext = consts.JPEG.String()
	case consts.SVG:
		ext = consts.SVG.String()
	default:
		ext = consts.PNG.String()
	}

	return path.Join(config.Share.File.Avatar, fmt.Sprintf(
		"%s-%d-%d.%s",
		req.GetHash(),
		req.GetSize(),
		time.Now().UnixNano(),
		ext,
	))
}

// 色值转换
func hexToRGBA(hexStr string) (color.Color, error) {
	hex := strings.ReplaceAll(hexStr, "#", "")

	if len(hex) != 6 {
		return nil, errors.New("非法颜色值")
	}

	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}, nil
}

// 默认填充色，从identicon包里扒出来的
func defaultFillColor(hashBytes []byte) color.Color {
	cb1, cb2 := uint32(hashBytes[0]), uint32(hashBytes[1])
	h := (cb1 + cb2) % 360
	s := (cb1 % 30) + 60
	l := (cb2 % 20) + 40

	if (h >= 50 && h <= 85) || (h >= 170 && h <= 190) {
		s = 80
		l -= 20
	} else if h > 85 && h < 170 {
		l -= 10
	}

	return identicon.HSL{H: h, S: s, L: l}
}

// 默认背景色，从identicon包里扒出来的
func defaultBackgroundColor() color.Color {
	return color.NRGBA{R: 240, G: 240, B: 240, A: 255}
}
