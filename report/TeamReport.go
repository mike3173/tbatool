package report

import (
	"fmt"
	"os"

	"github.com/mike3173/tbatool/models"
	"github.com/xuri/excelize/v2"
)

const WS_TEAMINFO = "team_info"

func TeamReport(team models.Team) {
	// Create new xlsx file
	var fname = fmt.Sprintf("%s-historical-performace.xlsx", team.Key)
	xlf := excelize.NewFile()

	// Create new worksheet team_info
	teamInfoSheet := xlf.NewSheet(WS_TEAMINFO)

	// Set Cell values
	xlf.SetCellValue(WS_TEAMINFO, "A1", "Team Info")
	xlf.SetCellValue(WS_TEAMINFO, "A2", "FRC Team:")
	xlf.SetCellValue(WS_TEAMINFO, "B2", fmt.Sprintf("%d %s", team.TeamNumber, team.Nickname))
	xlf.SetCellValue(WS_TEAMINFO, "A3", "Rookie Year:")
	xlf.SetCellValue(WS_TEAMINFO, "B3", fmt.Sprintf("%d", team.RookieYear))
	xlf.SetCellValue(WS_TEAMINFO, "A4", "From:")
	xlf.SetCellValue(WS_TEAMINFO, "B4", fmt.Sprintf("%s %s, %s %s", team.SchoolName, team.City, team.StateProv, team.Country))

	// Format Sheet1, named
	if err := xlf.MergeCell(WS_TEAMINFO, "A1", "B1"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create default Style
	defStyle, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 18},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defStyleBold, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 18},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defStyleBoldCenter, err := xlf.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 18},
		Alignment: &excelize.Alignment{
			Horizontal: "center"},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A1", "B4", defStyle); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A1", "A2", defStyleBoldCenter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetCellStyle(WS_TEAMINFO, "A2", "A4", defStyleBold); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetColWidth(WS_TEAMINFO, "A", "A", 19); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := xlf.SetColWidth(WS_TEAMINFO, "B", "B", 78); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	xlf.SetActiveSheet(teamInfoSheet)
	// Save and close
	if err := xlf.SaveAs(fname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
