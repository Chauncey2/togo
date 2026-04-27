package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

// 调色盘颜色坐标
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x振荡转几圈（影响曲线长途）
		res     = 0.001 // 角度步长（越小越精细）
		size    = 100   // 画布半径（最终宽高 2*size+1）
		nframes = 64    // 动画帧数
		delay   = 8     // 每帧间隔
	)
	freq := rand.Float64() * 3.0        //随机生成 y 振荡相对频率，范围约
	anim := gif.GIF{LoopCount: nframes} // 创建GIF动画对象
	phase := 0.0                        // 初始化相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 定义当前帧的矩形画布大小
		img := image.NewPaletted(rect, palette)      // 创建调色板图像
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 计算当前参数下的点坐标
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
