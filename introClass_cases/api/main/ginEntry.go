package main

import (
	"api/module"
	"api/utils/comMethods"
	"api/utils/log"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"path"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.POST("getDirInfo", getDirInfo)
	err := r.Run(":8081")
	if err != nil {
		log.Error().Msgf("service is not start,because:%s", err.Error())
	}
}

func getDirInfo(context *gin.Context) {
	const (
		baseDir = "/home/steed/Desktop/session_work/IntroClass"
	)
	log.Info().Msgf("projectName is : %s", context.PostForm("projectName"))
	projectName := context.PostForm("projectName")
	studentsInfo:=module.Info{}
	studentNames := comMethods.GetFileList(path.Join(baseDir, projectName))
	for _, stuName := range studentNames {
		var std module.Student
		jobNames := comMethods.GetFileList(path.Join(baseDir,projectName,stuName))
		for _, jn := range jobNames {
			std.AppendJob(jn)
			std.StuName = stuName
		}
		studentsInfo.AppendStudent(&std)
	}
	//log.Info().Msgf("info is ：%#v",studentsInfo)
	res,err:=json.Marshal(studentsInfo)

	if err != nil {
		log.Error().Msgf("marshal err is :%s",err.Error())
	}else {
		log.Info().Msgf("res is %s",string(res))
		context.JSON(200,gin.H{
			"status":"posted",
			"message":string(res),

		})

	}
}
