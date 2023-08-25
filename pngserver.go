 package main

 import (
 	"github.com/kbinani/screenshot"
 	//"image/png"
 	//"io"
    //"net/http"
    //"github.com/nfnt/resize"
    "github.com/gofiber/fiber/v2"
    //"fmt"
    "time"
    "image"
    "golang.org/x/image/draw"
    "encoding/json"
)





func main() {
    //images := make(chan image.RGBA, 8)

    app := fiber.New()

    app.Get("/batch/screenshot", func(c *fiber.Ctx) error {
        img, err := screenshot.CaptureDisplay(0)
        imgJson, _ := json.Marshal(image_2_array_pix(img))
        // c.Set("Content-Type", "image/png")
        c.SendString(string(imgJson))

        return err
    })

    go func() {
        for true {
            time.Sleep(time.Millisecond * 100)
        }
    }()

    app.Listen(":3000")
}

func image_2_array_pix(src image.Image) [][][3]float32 {
    bounds := src.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
    iaa := make([][][3]float32, height)
    src_rgba := image.NewRGBA(src.Bounds())
    draw.Copy(src_rgba, image.Point{}, src, src.Bounds(), draw.Src, nil)

    for y := 0; y < height; y++ {
        row := make([][3]float32, width)
        for x := 0; x < width; x++ {
            idx_s := (y * width + x) * 4
            pix := src_rgba.Pix[idx_s : idx_s + 4]
            row[x] = [3]float32{float32(pix[0]), float32(pix[1]), float32(pix[2])}
        }
        iaa[y] = row
    }

    return iaa
}
