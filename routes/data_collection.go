package routes

import (
	"dev-server/utils"
	"encoding/csv"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DataCollectionRequestData struct {
	Data [][]string `json:"Data"`
}

func data_collection(ctx *gin.Context) {
	var reqData DataCollectionRequestData

	err := ctx.BindJSON(&reqData)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
		})
	}

	// async write data to csv
	go func() {
		p := utils.GetCurrentAbPath()
		csvFilePath := filepath.Join(p, "../assets/data.csv")

		if !utils.Exist(csvFilePath) {
			createFile, err := os.Create(csvFilePath)
			if err != nil {
				panic(err)
			}
			defer createFile.Close()
		}

		txt, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer txt.Close()
		w := csv.NewWriter(txt) //创建一个新的写入文件流
		for _, v := range reqData.Data {
			w.Write(v)
			w.Flush()
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
