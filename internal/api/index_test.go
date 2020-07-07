package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/photoprism/photoprism/internal/i18n"
	"github.com/stretchr/testify/assert"
)

func TestCancelIndex(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		CancelIndexing(router)
		r := PerformRequest(app, "DELETE", "/api/v1/index")

		var resp i18n.Response

		if err := json.Unmarshal(r.Body.Bytes(), &resp); err != nil {
			t.Fatal(err)
		}

		assert.True(t, resp.Success())
		assert.Equal(t, i18n.Msg(i18n.MsgIndexingCanceled), resp.Msg)
		assert.Equal(t, i18n.Msg(i18n.MsgIndexingCanceled), resp.String())
		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
