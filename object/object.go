package object

import (
	"StrojProAPI/api"
	"StrojProAPI/app"
	"fmt"
	"github.com/CossackPyra/pyraconv"
	"github.com/tealeg/xlsx"
	"strings"
)

func Get(id string) Object {
	var i Object
	err := app.DB.Get(&i, `select * from objects where id=$1`, id)
	api.CheckErrInfo(err, "GetT")
	return i
}

func New(i Object) int {
	var id int
	err := app.DB.Get(&id, `INSERT INTO objects 
    (object_name, stage_name, sub_stage_name, period, percent_object, percent_stage) 
VALUES ($1,$2,$3,$4,$5,$6) returning id`,
		i.ObjectName, i.StageName, i.SubStageName, i.Period, i.PercentObject, i.PercentStage)
	api.CheckErrInfo(err, "NewT")
	return id
}

type Object struct {
	ID     int `json:"id,omitempty" db:"id,omitempty"`
	UserId int `json:"user_id,omitempty" db:"user_id,omitempty"`

	ObjectName   string `json:"object_name,omitempty" db:"object_name,omitempty"`
	StageName    string `json:"stage_name,omitempty" db:"stage_name,omitempty"`
	SubStageName string `json:"sub_stage_name,omitempty" db:"sub_stage_name,omitempty"`

	Period      int `json:"period,omitempty" db:"period,omitempty"`
	TotalPeriod int `json:"total_period,omitempty" db:"total_period,omitempty"`

	PercentObject     float64 `json:"percent_object,omitempty" db:"percent_object,omitempty"`
	PercentStage      float64 `json:"percent_stage,omitempty" db:"percent_stage,omitempty"`
	TotalPercentStage float64 `json:"total_percent_stage,omitempty" db:"total_percent_stage,omitempty"`

	IsCompleted   bool        `json:"is_completed,omitempty" db:"is_completed,omitempty"`
	CompletedTime interface{} `json:"completed_time,omitempty" db:"completed_time,omitempty"`

	TotalPercent float64 `json:"total_percent,omitempty" db:"total_percent,omitempty"`
}

func ParseFILE() {
	excelFileName := "C:\\Users\\Danii\\Downloads\\Telegram Desktop\\Таблица.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("err file xlsx")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var t Object
			t.ObjectName = pyraconv.ToString(row.Cells[0])
			t.StageName = pyraconv.ToString(row.Cells[1])
			t.SubStageName = pyraconv.ToString(row.Cells[2])

			t.Period = int(pyraconv.ToFloat64(row.Cells[3]))

			t.PercentObject = pyraconv.ToFloat64(strings.TrimSuffix(api.ToString(row.Cells[4]), "%"))

			t.PercentStage = pyraconv.ToFloat64(strings.TrimSuffix(api.ToString(row.Cells[5]), "%"))
			fmt.Println(New(t))
		}
	}
}
