package api

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"server-application/models"
	"server-application/server"

	"github.com/disintegration/imaging"
	uuid "github.com/gofrs/uuid"
)

var files = new(apiFiles)

type apiFiles struct{}

func handleFiles(r *server.Router) {
	r.GET("/{name}", files.show)
	r.POST("", files.upload)
	r.GET("", files.list)
}

type FilesResponse struct {
	ID   int               `json:"id"`
	URLs map[string]string `json:"urls"`
}

type ImageResizeParams struct {
	Width   int
	Height  int
	Quality int
}

var ImageParams = map[string]ImageResizeParams{
	"original": {0, 0, 0},
	"800x800":  {800, 800, 80},
	"400x400":  {400, 400, 80},
	"150x150":  {150, 150, 80},
	"50x50":    {50, 50, 80},
}

func getURLs(dir string) map[string]string {
	urls := map[string]string{}
	for k := range ImageParams {
		urls[k] = fmt.Sprintf("http://127.0.0.1/img/%s/%s.jpg", dir, k)
	}

	return urls
}

func (apiFiles) response(f *models.File) *FilesResponse {
	return &FilesResponse{
		ID:   f.ID,
		URLs: getURLs(f.Name),
	}
}

func (af *apiFiles) multiResponse(fs []*models.File) []*FilesResponse {
	fr := []*FilesResponse{}
	for _, f := range fs {

		resp := af.response(f)
		if resp != nil {
			fr = append(fr, resp)
		}

	}

	return fr
}

type HTTPfile interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

func resizeAndStore(dir string, file HTTPfile, name string, params ImageResizeParams) error {
	if name == "original" {
		file.Seek(0, 0)
		buf := &bytes.Buffer{}
		_, err := buf.ReadFrom(file)
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("assets/img/%s/%s.jpg", dir, name), buf.Bytes(), 0o777)
		if err != nil {
			return err
		} else {
			file.Seek(0, 0)
			img, _, err := image.Decode(file)
			if err != nil {
				return err
			}
			dImg := imaging.Fit(img, params.Width, params.Height, imaging.Lanczos)
			var b []byte
			buf := bytes.NewBuffer(b)
			jpeg.Encode(buf, dImg, &jpeg.Options{Quality: params.Quality})
			err = os.WriteFile(fmt.Sprintf("assets/img/%s/%s.jpg", dir, name), buf.Bytes(), 0o777)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func (af *apiFiles) upload(c *server.Context) {
	HTTPFile, _, err := c.Request.FormFile("file")
	if err != nil {
		c.RenderError(http.StatusBadRequest, err)
		return

	}

	//	dir := uuid.NewV1().String()
	id, err := uuid.NewV1()
	if err != nil {
		c.RenderError(http.StatusInternalServerError, err)
		return
	}
	dir := id.String()

	err = os.Mkdir(fmt.Sprintf("assets/img/%s", dir), 0o777)
	if err != nil {
		c.RenderError(http.StatusBadRequest, err)
		return

	}

	for k, p := range ImageParams {
		err = resizeAndStore(dir, HTTPFile, k, p)
		if err != nil {
			c.RenderError(http.StatusBadRequest, err)
			return
		}

	}

	file, err := models.Files.Create(dir)
	if err != nil {
		c.RenderError(http.StatusBadRequest, err)
		return

	}

	c.RenderJSON(http.StatusCreated, af.response(file))
}

func (af *apiFiles) show(c *server.Context) {
	file, err := models.Files.ByName(c.Param("name"))
	if err != nil {
		c.RenderError(http.StatusBadRequest, err)
		return

	}
	c.RenderJSON(http.StatusOK, af.response(file))
}

func (af *apiFiles) list(c *server.Context) {
	files, err := models.Files.List("")
	if err != nil {
		c.RenderError(http.StatusBadRequest, err)
		return

	}

	c.RenderJSON(http.StatusOK, af.multiResponse(files))
}
