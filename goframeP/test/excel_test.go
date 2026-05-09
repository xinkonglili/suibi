package test

import (
	"fmt"
	"math"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestExists(t *testing.T) {
	f, err := excelize.OpenFile("C:\\Users\\admin\\Desktop\\网约车订单列表.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sheet := "Sheet1"
	rows, err := f.GetRows(sheet)
	if err != nil {
		panic(err)
	}

	header := rows[0]
	aaIdx, ahIdx, orderIdx, iIdx, fIdx, dIdx := -1, -1, -1, -1, -1, -1
	for i, col := range header {
		if col == "实际接单下游价格（元）" {
			aaIdx = i
		}
		if col == "原始行程费总计（元）" {
			ahIdx = i
		}
		if col == "渠道主单号" {
			orderIdx = i
		}
		if i == 8 { // I列 index=8
			iIdx = i
		}
		if i == 5 { // F列 index=5
			fIdx = i
		}
		if i == 3 { // D列 index=3
			dIdx = i
		}
	}
	if aaIdx == -1 || ahIdx == -1 {
		panic("找不到 AA 或 AH 列")
	}

	out := excelize.NewFile()
	outSheet := "Sheet1"
	out.SetSheetRow(outSheet, "A1", &header)

	resultRow := 2
	for i, row := range rows[1:] {
		if aaIdx >= len(row) || ahIdx >= len(row) {
			continue
		}
		var aa, ah float64
		fmt.Sscanf(row[aaIdx], "%f", &aa)
		fmt.Sscanf(row[ahIdx], "%f", &ah)

		if math.Round(aa*100) != math.Round(ah*100) {
			cell, _ := excelize.CoordinatesToCellName(1, resultRow)
			out.SetSheetRow(outSheet, cell, &row)
			resultRow++

			orderNo := ""
			if orderIdx != -1 && orderIdx < len(row) {
				orderNo = row[orderIdx]
			}
			iVal := ""
			if iIdx != -1 && iIdx < len(row) {
				iVal = row[iIdx]
			}
			fVal := ""
			if fIdx != -1 && fIdx < len(row) {
				fVal = row[fIdx]
			}
			dVal := ""
			if dIdx != -1 && dIdx < len(row) {
				dVal = row[dIdx]
			}
			fmt.Printf("第 %d 行不相等: 渠道主单号=%s, Switch主单号=%s, Switch下游子单号=%s, 运力子单号=%s, AA=%.2f, AH=%.2f\n", i+2, orderNo, dVal, fVal, iVal, aa, ah)
		}
	}

	out.SaveAs("筛选结果.xlsx")
	fmt.Printf("\n共筛选出 %d 行，已保存到 筛选结果.xlsx\n", resultRow-2)
}
