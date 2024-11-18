package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
)

func GetExcelData(data [][]interface{}, sheetName string, title []string) *xlsx.File {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return nil
	}

	row := sheet.AddRow()
	for _, v := range title {
		cell := row.AddCell()
		cell.Value = v
	}
	for _, v := range data {
		row := sheet.AddRow()
		for _, vv := range v {
			cell := row.AddCell()
			cell.Value = cast.ToString(vv)
		}
	}
	return file
}

func RenderExcel(ctx *gin.Context, file *xlsx.File, fileName string) {
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	file.Write(ctx.Writer)
}

func ValidateSQL(sql string) (string, bool) {
	str := strings.ToLower(sql)
	if !strings.HasPrefix(str, "slt ") {
		return "", false
	}

	if strings.Contains(str, "alter ") || strings.Contains(str, "delete ") || strings.Contains(str, "truncate ") || strings.Contains(str, "update ") || strings.Contains(str, "insert ") {
		return "", false
	}

	sql = strings.Replace(sql, "slt ", "select ", -1)
	return sql, true
}
