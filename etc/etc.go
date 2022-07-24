package etc

import (
	"StrojProAPI/api"
	"StrojProAPI/app"
	"StrojProAPI/object"
)

func ExistObjectName(i string) bool {
	var exist bool
	err := app.DB.Get(&exist, `SELECT exists(select id FROM objects where  object_name = $1);`, i)
	api.CheckErrInfo(err, "ExistObjectName")
	return exist
}

func ExistStageName(i string) bool {
	var exist bool
	err := app.DB.Get(&exist, `SELECT exists(select id FROM objects where  stage_name = $1);`, i)
	api.CheckErrInfo(err, "ExistStageName")
	return exist
}

func ExistSubStageName(i string) bool {
	var exist bool
	err := app.DB.Get(&exist, `SELECT exists(select id FROM objects where  sub_stage_name = $1);`, i)
	api.CheckErrInfo(err, "ExistSubStageName")
	return exist
}

func GetStagesByObjName(objName string) []object.Object {
	rows, err := app.DB.Queryx("SELECT  stage_name, sum(percent_stage) as total_percent_stage, sum(period) as total_period FROM objects where object_name=$1 group by stage_name, is_completed;", objName)
	api.CheckErrInfo(err, "GetStagesByObjName 1")

	i := []object.Object{}

	for rows.Next() {
		var item object.Object
		err = rows.StructScan(&item)
		api.CheckErrInfo(err, "GetStagesByObjName 2")
		i = append(i, item)
	}
	_ = rows.Close()
	return i
}

func GetObjList() []object.Object {
	rows, err := app.DB.Queryx("SELECT  object_name, sum(percent_object) as total_percent FROM objects group by object_name;")
	api.CheckErrInfo(err, "GetObjList 1")

	i := []object.Object{}

	for rows.Next() {
		var item object.Object
		err = rows.StructScan(&item)
		api.CheckErrInfo(err, "GetObjList 2")
		i = append(i, item)
	}
	_ = rows.Close()
	return i
}

func GetSubStageListByObjNameStage(objName, stageName string) []object.Object {
	rows, err := app.DB.Queryx("SELECT  sub_stage_name, percent_stage, period, is_completed, completed_time FROM objects where object_name=$1 and stage_name=$2 order by id;", objName, stageName)
	api.CheckErrInfo(err, "GetSubStageListByObjNameStage 1")

	i := []object.Object{}

	for rows.Next() {
		var item object.Object
		err = rows.StructScan(&item)
		api.CheckErrInfo(err, "GetSubStageListByObjNameStage 2")
		i = append(i, item)
	}
	_ = rows.Close()
	return i
}
