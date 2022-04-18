package http_service

import (
	"fmt"
	"stream-api/models"
)

// Gets Stream URL from Betradar
func GetStreamUrl(streamId, streamType, clientIp string) (models.Response, error) {
	bs := BetradarHttpService{}

	return bs.Get(fmt.Sprintf("streams/%s/%s", streamId, streamType), clientIp)
}
