package main

import (
	"api/module"
	"api/utils/comMethods"
	"api/utils/log"
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
		baseDir = "/home/steed/Documents/myWork/case_generate/IntroClass"
	)
	log.Info().Msgf("projectName is : %s", context.PostForm("projectName"))
	projectName := context.PostForm("projectName")
	//if strings.TrimSpace(projectName) == ""{
	//	log.Error().Msgf("get empty key word from front end!")
	//	panic(errors.New("http data flow error"))
	//}
	studentsInfo:=module.Info{}
	studentNames := comMethods.GetFileList(path.Join(baseDir, projectName))
	for _, stuName := range studentNames {
		var std module.Student
		jobNames := comMethods.GetFileList(path.Join(baseDir,projectName,stuName))
		for _, jn := range jobNames {
			//log.Info().Msgf("jobname is %s",jn)
			std.AppendJob(jn)
			std.StuName = stuName
		}
		studentsInfo.AppendStudent(&std)
	}
	//log.Info().Msgf("info is ：%#v",studentsInfo)
	//res,err:=json.Marshal(studentsInfo)
	context.PureJSON(200,gin.H{
		"status":"posted",
		"message":studentsInfo,

	})
	//if err != nil {
	//	log.Error().Msgf("marshal err is :%s",err.Error())
	//}else {
	//	log.Info().Msgf("res is %#v",string(res))
	//	context.PureJSON(200,gin.H{
	//		"status":"posted",
	//		"message":studentsInfo,
	//
	//	})

	//}
}
