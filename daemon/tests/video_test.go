package test

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"rkl.io/kika-downloader/daemon/dto"
	"github.com/stretchr/testify/assert"
)

// TestGetVideo test to get a single video from an URL
func TestGetVideo(t *testing.T) {
	videoURL := base64.URLEncoding.EncodeToString([]byte("http://www.kika.de/super-wings/sendungen/sendung97844.html"))

	r, err := http.NewRequest("GET", "/api/v1/video/"+videoURL, nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetVideo", "Code[%d]\n%s", w.Code, w.Body.String())

	var videoDto dto.Video

	if err := json.Unmarshal(w.Body.Bytes(), &videoDto); err != nil {
		t.Error(err)
	}

	a := assert.New(t)

	a.Equal(videoDto.SeriesTitle, "Super Wings", "expected \"Super Wings\" as series title")
	a.Equal(videoDto.EpisodeTitle, "Schwingt die Pinsel!", "expected \"Schwingt die Pinsel!\" as episode title")
	a.Equal(videoDto.DownloadURL, "http://pmdonline.kika.de/mp4dyn/7/FCMS-7bdd1753-dd4d-404e-a652-d94ec2183363-5a2c8da1cdb7_7b.mp4", "wrong download url")
	a.Equal(videoDto.EpisodeNo, 23, "expected \"23\" as episode number")
}
