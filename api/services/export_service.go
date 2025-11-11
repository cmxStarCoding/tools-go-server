package services

import (
	"encoding/csv"
	"fmt"
	"journey/models"
	"os"
	"strings"
	"time"

	"journey/common/database"
	"journey/common/pkg"

	"github.com/xuri/excelize/v2"
)

type ExportService struct{}

func (s ExportService) AsyncExport() {
	var exportTask models.TExportTask
	err := database.DB.Where("status = ?", 1).First(&exportTask).Error
	if err != nil {
		fmt.Println("没有待处理任务或查询出错：", err)
		return
	}

	fmt.Println("开始处理任务：", exportTask.ID)

	sqlText := strings.TrimSpace(exportTask.SqlText)
	if !strings.HasPrefix(strings.ToLower(sqlText), "select") {
		s.updateTaskFail(&exportTask, "仅支持SELECT语句")
		return
	}

	// === 1️⃣ 统计总数 ===
	countSQL := strings.TrimSpace(exportTask.CountSqlText)
	var total int64
	if err := database.DB.Raw(countSQL).Scan(&total).Error; err != nil {
		s.updateTaskFail(&exportTask, "统计总数失败："+err.Error())
		return
	}

	if total == 0 {
		s.updateTaskFail(&exportTask, "没有可导出的数据")
		return
	}

	fmt.Printf("共需导出 %d 条记录\n", total)

	// === 2️⃣ 判断导出类型 ===
	isCSV := total > 10000

	fileDir := "exports"
	_ = os.MkdirAll(fileDir, 0755)
	ext := "xlsx"
	if isCSV {
		ext = "csv"
	}
	filePath := fmt.Sprintf("%s/%d_%s.%s", fileDir, exportTask.ID, time.Now().Format("20060102150405"), ext)

	// === 3️⃣ 创建导出配置 ===
	exporter := pkg.NewExport(
		map[string]string{
			"id":         "用户ID",
			"phone":      "用户手机号码",
			"account":    "用户账号",
			"nickname":   "用户昵称",
			"avatar_url": "用户头像",
			"created_at": "创建时间",
		},
		map[string]int{
			"A": 10,
			"B": 12,
			"C": 20,
			"D": 10,
			"E": 20,
			"F": 20,
		},
	)

	pageSize := 10
	if isCSV {
		err = s.exportCSV(sqlText, filePath, int(total), pageSize)
	} else {
		err = s.exportExcel(sqlText, filePath, int(total), pageSize, exporter)
	}

	if err != nil {
		s.updateTaskFail(&exportTask, err.Error())
		return
	}

	exportTask.Status = 2
	exportTask.FilePath = filePath
	database.DB.Save(&exportTask)
	fmt.Println("✅ 导出完成，文件路径：", filePath)
}

// CSV 导出
func (s ExportService) exportCSV(sqlText, filePath string, total, pageSize int) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建CSV文件失败: %v", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	isHeaderWritten := false

	for offset := 0; offset < total; offset += pageSize {
		pageSQL := fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlText, pageSize, offset)
		rows, err := database.DB.Raw(pageSQL).Rows()
		if err != nil {
			return fmt.Errorf("查询失败: %v", err)
		}
		defer rows.Close()

		cols, _ := rows.Columns()
		if !isHeaderWritten {
			writer.Write(cols)
			isHeaderWritten = true
		}

		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}
			rows.Scan(columnPointers...)

			record := make([]string, len(cols))
			for i, col := range columns {
				if b, ok := col.([]byte); ok {
					record[i] = string(b)
				} else if col == nil {
					record[i] = ""
				} else {
					record[i] = fmt.Sprintf("%v", col)
				}
			}
			writer.Write(record)
		}
		writer.Flush()
		rows.Close()
	}

	return nil
}

// Excel 导出（支持列宽、表头映射）
func (s ExportService) exportExcel(sqlText, filePath string, total, pageSize int, exporter *pkg.Export) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)

	rowIndex := 1
	isHeaderWritten := false

	for offset := 0; offset < total; offset += pageSize {
		pageSQL := fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlText, pageSize, offset)
		rows, err := database.DB.Raw(pageSQL).Rows()
		if err != nil {
			return fmt.Errorf("查询失败: %v", err)
		}
		defer rows.Close()

		cols, _ := rows.Columns()
		if !isHeaderWritten {
			for i, col := range cols {
				header := col
				if alias, ok := exporter.Headers[col]; ok {
					header = alias
				}
				cell, _ := excelize.CoordinatesToCellName(i+1, rowIndex)
				f.SetCellValue(sheet, cell, header)
			}

			// 设置列宽
			for col, width := range exporter.ColumnWidth {
				f.SetColWidth(sheet, col, col, float64(width))
			}
			rowIndex++
			isHeaderWritten = true
		}

		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}
			rows.Scan(columnPointers...)

			for i, col := range columns {
				val := ""
				if b, ok := col.([]byte); ok {
					val = string(b)
				} else if col != nil {
					val = fmt.Sprintf("%v", col)
				}
				cell, _ := excelize.CoordinatesToCellName(i+1, rowIndex)
				f.SetCellValue(sheet, cell, val)
			}
			rowIndex++
		}
		rows.Close()
	}

	if err := f.SaveAs(filePath); err != nil {
		return fmt.Errorf("保存Excel失败: %v", err)
	}
	return nil
}

// 更新任务状态
func (s ExportService) updateTaskFail(task *models.TExportTask, reason string) {
	task.Status = 3
	task.FailReason = reason
	database.DB.Save(task)
}
