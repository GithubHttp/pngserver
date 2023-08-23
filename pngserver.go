 package main

 import (
 	"github.com/kbinani/screenshot"
 	"image/png"
 	"io"
    "net/http"
    "github.com/nfnt/resize"
 )

 func saveScreenshot(w io.Writer, width uint, height uint) {
     bounds := screenshot.GetDisplayBounds(0)

     img, err := screenshot.CaptureRect(bounds)
     if err != nil {
         panic(err)
     }
     resizedImg := resize.Resize(width, height, img, resize.Bilinear)
     png.Encode(w, resizedImg)
}

func screenshotHandler(w http.ResponseWriter, r *http.Request) {
    saveScreenshot(w, 1280, 720)
}

func main() {
    http.HandleFunc("/screenshot", screenshotHandler)

    http.ListenAndServe(":80", nil)
}
