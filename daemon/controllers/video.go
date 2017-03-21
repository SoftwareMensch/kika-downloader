package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/sarulabs/di"
	"rkl.io/kika-downloader/core/config"
	"rkl.io/kika-downloader/core/contract"
	"rkl.io/kika-downloader/daemon/dto"
)

// VideoController operations for Video
type VideoController struct {
	beego.Controller

	coreContext di.Context
}

// URLMapping ...
func (c *VideoController) URLMapping() {
	c.Mapping("Get", c.Get)
}

func (c *VideoController) ServeJsonException(code int16, message string) {
	c.Ctx.Output.SetStatus(int(code))
	c.Data["json"] = dto.NewException(code, message)
	c.ServeJSON()
	c.Abort(string(code))
}

func (c *VideoController) Prepare() {
	coreBuilder, err := config.InitCoreBuilder("socks5://192.168.247.1:9050")
	if err != nil {
		c.ServeJsonException(500, err.Error())
	}

	c.coreContext = coreBuilder.Build()
}

// @Title Get
// @Description get video by video url
// @Param	url		path 	string	true		"the video url you want to get"
// @Success 200 {object} models.Video
// @Failure 403 :objectId is empty
// @router /:url [get]
func (c *VideoController) Get() {
	encodedURL := c.Ctx.Input.Param(":url")

	decodedURLData, err := base64.URLEncoding.DecodeString(encodedURL)
	if err != nil {
		c.ServeJsonException(400, "malformed encoded source url")
	}

	decodedURL := string(decodedURLData)

	service, err := c.coreContext.SafeGet("video_extractor")
	if err != nil {
		c.ServeJsonException(500, "video extractor service not found")
	}
	videoExtractor := service.(contract.VideoExtractorInterface)

	video, err := videoExtractor.ExtractVideoFromURL(decodedURL)
	if err != nil {
		c.ServeJsonException(500, err.Error())
	}

	videoDto := dto.NewVideoDtoFromModel(video)

	c.Data["json"] = &videoDto
	c.ServeJSON()
}
