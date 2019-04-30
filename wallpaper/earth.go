package wallpaper

// 拼接图片地址示例: http://himawari8-dl.nict.go.jp/himawari8/img/D531106/4d/550/2016/01/08/035000_0_0.png
// 其中的`4d`表示 4 倍尺寸的图片. 可选 1d, 2d, 4d, 8d, 16d.

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"net/http"
	"os"
	"regexp"

	"go-wallpaper/utils"

	"github.com/astaxie/beego"
	"github.com/reujab/wallpaper"
)

var workDir = fmt.Sprintf("%s/earth", beego.AppConfig.String("WallpaperPath"))

func EarthRun() {
	if err := os.MkdirAll(workDir, 0755); err != nil {
		beego.Error(err)
	}

	// EarthIncreaseLevel:
	// Increases the quality and the size. Possible values: 4, 8, 16, 20
	level, _ := beego.AppConfig.Int("earth::EarthIncreaseLevel")
	w := NewWallpaper(level)
	if err := w.cacheImages(); err != nil {
		beego.Error(err)
	}
	if err := w.concat(); err != nil {
		beego.Error(err)
	}

	if err := wallpaper.SetFromFile(w.Path); err != nil {
		beego.Error(err)
	}
}

type Wallpaper struct {
	Multiple  int
	Prefix    string
	PartsPath [][]string
	Path      string
}

func NewWallpaper(multiple int) *Wallpaper {
	latestDate, err := getLatestDate()
	if err != nil {
		beego.Error(err)
	}
	return &Wallpaper{
		Multiple: multiple,
		Prefix:   formatImagePrefix(latestDate),
		Path:     fmt.Sprintf("%s/final_%s.png", workDir, latestDate[14:16]),
	}
}

func (w *Wallpaper) cacheImages() error {
	paths := make([][]string, w.Multiple)
	for r := 0; r < w.Multiple; r++ {
		paths[r] = make([]string, w.Multiple)
		for c := 0; c < w.Multiple; c++ {
			url := fmt.Sprintf("https://himawari8-dl.nict.go.jp/himawari8/img/D531106/%dd/550/%s_%d_%d.png",
				w.Multiple, w.Prefix, c, r)
			path := fmt.Sprintf("%s/part_%d_%d.png", workDir, c, r)
			if err := saveFile(url, path); err != nil {
				return err
			}
			fmt.Println("cache", path)
			paths[r][c] = path
		}
	}
	w.PartsPath = paths
	return nil
}

func (w *Wallpaper) concat() error {
	L := 550

	img := image.NewRGBA(image.Rect(0, 0,
		L*w.Multiple,
		L*w.Multiple,
	))

	for r := 0; r < w.Multiple; r++ {
		for c := 0; c < w.Multiple; c++ {
			f, err := os.Open(w.PartsPath[r][c])
			if err != nil {
				return err
			}
			part, err := png.Decode(f)
			f.Close()
			if err != nil {
				return err
			}
			draw.Draw(
				img,
				part.Bounds().Add(image.Pt(L*c, L*r)),
				part,
				part.Bounds().Min,
				draw.Over,
			)
		}
	}

	file, err := os.Create(w.Path)
	if err != nil {
		return err
	}
	png.Encode(file, img)
	file.Close()
	beego.Info("Wallpaper saved:", w.Path)
	return nil
}

// return example: "2018/05/16/235000"
func formatImagePrefix(latestDate string) string {
	s := latestDate
	s = regexp.MustCompile(`[-, ]`).ReplaceAllString(s, "/")
	s = regexp.MustCompile(`:`).ReplaceAllString(s, "")
	return s
}

// return example: "2018-05-16 23:50:00"
func getLatestDate() (date string, err error) {
	var payload struct {
		Date string `json:"date"`
	}
	_, err = utils.GetJSON(&payload, "https://himawari8-dl.nict.go.jp/himawari8/img/D531106/latest.json", nil)
	if err != nil {
		return "", err
	}
	return payload.Date, nil
}

func saveFile(url string, path string) error {
	rsp, err := http.Get(url)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()
	defer f.Close()

	_, err = io.Copy(f, rsp.Body)
	if err != nil {
		return err
	}
	return nil
}
