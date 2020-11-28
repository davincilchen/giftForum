package ginprocess

import (



	"github.com/gin-gonic/gin"
)

func GetRawData(ctx *gin.Context) ([]byte, error) {

	body, err := ctx.GetRawData()
	if err != nil {
		return nil, err
	}

	return body, nil
}
