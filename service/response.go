package service

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Meta map[string]interface{} `json:"meta"`
	Data interface{}            `json:"data"`
}

func Response(data interface{}, c *gin.Context, table string, total int64) interface{} {
	var metaData map[string]interface{}

	// if table != "" {
	// 	metaData = make(map[string]interface{})
	// 	perPage, _ := strconv.Atoi(c.URL.Query().Get("perPage"))
	// 	if perPage == 0 {
	// 		perPage = 10
	// 	}
	// 	page, _ := strconv.Atoi(c.URL.Query().Get("page"))
	// 	if page == 0 {
	// 		page = 1
	// 	}

	// 	var last_page int64

	// 	for i := int64(0); i < total; i++ {
	// 		if i%int64(perPage) == 0 {
	// 			last_page++
	// 		}
	// 	}
	// 	to := int64(page) * int64(perPage)
	// 	if int64(page) == last_page {
	// 		to = int64(page)*int64(perPage) - (int64(page)*int64(perPage) - total)
	// 	} else if total == 0 {
	// 		to = 0
	// 	}

	// 	metaData["total"] = total
	// 	metaData["per_page"] = perPage
	// 	metaData["current_page"] = page
	// 	metaData["from"] = page*perPage - perPage + 1
	// 	metaData["to"] = to
	// 	metaData["last_page"] = last_page
	// }

	result := response{
		Meta: metaData,
		Data: data,
	}

	return result
}
