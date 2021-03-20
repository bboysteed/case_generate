<template>
  <div>
    1.选择学生提交的作业：
    <el-cascader
        style="margin-top: 20px"
        class="el-cascader--medium"
        expand-trigger="hover"
        :options="jobOptions"
        v-model="jobSelectedOptions"
        @change="handleChange">
    </el-cascader>
    2.选择算法：
    <el-cascader
        style="margin-top: 20px"
        class="el-cascader--medium"
        expand-trigger="hover"
        :options="algorithmOptions"
        v-model="algoSelectedOptions"
        @change="handleAlgoChange">
    </el-cascader>
    <el-button icon="el-icon-check" round type="primary" style="margin-left: 30px" @click="commitMission">
      生成测试用例
    </el-button>
    <el-col :span="12" style="margin: 30px">
      <codemirror :value="codeContent" :options="cmOptions" class="code-mirror"></codemirror>
    </el-col>

    <el-col :span="12" class="el-col-push-3" style="margin-top: 20px">
      <div slot="header" class="clearfix">
        <span>用例列表</span>
      </div>
      <ul class="infinite-list" v-infinite-scroll="load" style="overflow:auto">
        <li v-for="(acase,indx) in caseContent" :key="indx" class="infinite-list-item">({{ indx }})--- {{ acase }}</li>
      </ul>

    </el-col>

  </div>
</template>

<script>
// import selectOptions from '@/statics/js/options'

import 'codemirror/lib/codemirror.css'
import 'codemirror/mode/clike/clike.js'

export default {

  name: "CaseGenerate",
  data() {
    return {
      projectName: null,
      jobOptions: null,
      jobSelectedOptions: ['资源', '组件交互文档'],
      algoSelectedOptions: ['遗传算法'],
      algorithmOptions: [
        {
          value: "GA",
          label: "遗传算法",
          children: [
            {
              value: "TSP",
              label: "旅行商遗传算法",
            }, {
              value: "DA",
              label: "差分进化算法",
            }
          ]
        }, {
          value: "",
          label: "神经网络",
          children: [
            {
              value: "",
              label: "LSTM",
            }, {
              value: "",
              label: "等等",
            }
          ]
        }
      ],
      codeContent: '',
      caseContent: ["1,2,3", "4,5,6"],
      loading: false,
      cmOptions: {
        lineNumbers: true,
        matchBrackets: true,
        readOnly: true,
        mode: 'text/x-c',
        tabSize: 4,
        line: true,
        theme:"3024-night"
      }

    };
  },
  mounted() {
    this.projectName = this.$route.params.projectName
    this.getSelectMenu()

  },
  methods: {
    handleAlgoChange(value) {
      if (value[1] !== "TSP") {
        this.$message.error("目前仅支持TSP遗传算法")
      } else {
        this.$message.success("success")
      }
    },
    handleChange(value) {
      this.$axios
          .get("/getCodeContent", {
            params: {
              projectName: this.projectName,
              stuName: value[0],
              jobName: value[1],
            }
          })
          .then((response) => {
            if (response.status === 400) {
              console.log("失败")
              this.$message.error("文件占用")
            } else if (response.status === 200) {
              this.$message.success("success")
              this.codeContent = response.data.codeContent
              // console.log(this.codeContent)
            }
          })
          .catch((error) => {
            console.log(error)
          })
    },
    getSelectMenu() {
      this.$axios
          .post("/getDirInfo", this.$qs.stringify(
              {
                "projectName": this.projectName
              }
          ))
          .then((response) => {
            this.jobOptions = response.data.message.all_students
            // console.log(this.stuInfo)
          })
          .catch((error) => {
            console.log(error)
          })
    },
    commitMission() {
      this.loading = true
      this.$axios
          .get("/geneCases", {
            params: {
              projectName: this.projectName,
            }
          })
          .then((response) => {
            console.log(response)

            if (response.data.status==="submit ok") {
              this.$message.success("提交成功")
            } else if (response.data.status==="submit error") {
              this.$message.error("提交失败")
            }
          })
          .catch((error) => {
            console.log(error)
          })
      this.loading = false
    }
  }

}
</script>

<style>
.CodeMirror {
  height: 600px;
  /*display: inline-block;*/
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  font-size: 15px;
  float: left;
}

.text {
  font-size: 14px;
}

.item {
  padding: 2px 0;
}


</style>