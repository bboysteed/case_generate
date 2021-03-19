package main

import (
	"api/module"
	"api/utils/comMethods"
	"api/utils/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path"
	"strings"
)

const (
	introClassBaseDir        = "/home/steed/Desktop/session_work/git_work/case_generate/IntroClass"
	introClassCaseGenBaseDir = "/home/steed/Desktop/session_work/git_work/case_generate/introClass_cases"
	pythonDir                = "/bin/python"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.POST("/getDirInfo", getDirInfo)
	r.GET("/getCodeContent", getCodeContent)
	r.GET("/geneCases", geneCases)
	err := r.Run(":8081")
	if err != nil {
		log.Error().Msgf("service is not start,because:%s", err.Error())
	}
}

func geneCases(context *gin.Context) {
	projectName := context.Query("projectName")
	log.Info().Msg( path.Join(introClassCaseGenBaseDir,projectName,"gen_cases.py"))
	_, err := exec.Command(pythonDir, path.Join(introClassCaseGenBaseDir,projectName,"gen_cases.py")).Output()
	if err != nil {
		log.Error().Msgf("run script failed ,err is:%s",err.Error())
		context.JSON(http.StatusInternalServerError,gin.H{
			"status":"fail to start gen case script",
			"reason":"not sure",
			"cases":"",
		})
	}else {
		buf,err:=ioutil.ReadFile(path.Join(introClassCaseGenBaseDir,"temFile","cases"))
		if err != nil {
			log.Error().Msgf("read file err:%s",err.Error())
			return
		}else {
			cases := strings.Split(string(buf),"\n")
			context.JSON(http.StatusOK,gin.H{
				"status":"ok",
				"reason":"",
				"cases":cases,
			})
		}

	}
}

func getCodeContent(context *gin.Context) {
	projectName := context.Query("projectName")
	stuName := context.Query("stuName")
	jobName := context.Query("jobName")
	log.Info().Msgf("file path is :%s", path.Join(introClassBaseDir, projectName, stuName, jobName))
	content, err := ioutil.ReadFile(path.Join(introClassBaseDir, projectName, stuName, jobName, projectName+".c"))
	log.Info().Msgf("file content is :%s", string(content))
	if err != nil {
		log.Info().Msgf("err open file,err is:%s", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":      "can't get file now",
			"reason":      "busy",
			"codeContent": "",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"reason":      "",
			"codeContent": string(content),
		})
	}
}

func getDirInfo(context *gin.Context) {

	log.Info().Msgf("projectName is : %s", context.PostForm("projectName"))
	projectName := context.PostForm("projectName")
	studentsInfo := module.Info{}
	studentNames := comMethods.GetFileList(path.Join(introClassBaseDir, projectName))
	for _, stuName := range studentNames {
		var std module.Student
		jobNames := comMethods.GetFileList(path.Join(introClassBaseDir, projectName, stuName))
		for _, jn := range jobNames {
			std.AppendJob(jn, jn)
			std.StuName = "stu--" + stuName[:8]
			std.StuValue = stuName
		}
		studentsInfo.AppendStudent(&std)
	}
	log.Info().Msg(studentsInfo.AllStudents[0].StuValue)
	context.JSON(200, gin.H{
		"status":  "posted",
		"message": &studentsInfo,
	})

}
