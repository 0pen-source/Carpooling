package common

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/xxtea/xxtea-go/xxtea"
)

var _ render.Render = EncryptedJSON{}

// EncryptedJSON is a render that encrypts data
type EncryptedJSON struct {
	render.JSON
	Key []byte
}

// Render _
func (r EncryptedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	data, err := json.Marshal(r.Data)
	if err != nil {
		panic(err)
	}

	if _, err = w.Write(xxtea.Encrypt(data, r.Key)); err != nil {
		panic(err)
	}

	return nil
}

// NewEncryptedJSONRender creates a new encrypted json render
func NewEncryptedJSONRender(data interface{}, key []byte) render.Render {
	r := EncryptedJSON{}
	r.Data = data
	r.Key = key
	return r
}
