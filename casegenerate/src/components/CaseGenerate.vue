<template>

  <div>
    <div>
      <span style="">1.选择学生提交的作业：</span>
      <el-cascader
          style="margin-top: 20px"
          class="el-cascader--medium"
          expand-trigger="hover"
          :options="jobOptions"
          v-model="jobSelectedOptions"
          @change="handleChange">
      </el-cascader>
      <span style="margin-left: 20px">2.选择算法：</span>
      <el-cascader
          style="margin-top: 20px;"
          class="el-cascader--medium"
          expand-trigger="hover"
          :options="algorithmOptions"
          v-model="algoSelectedOptions"
          @change="handleAlgoChange">
      </el-cascader>
      <el-button icon="el-icon-check" round type="primary" style="margin-top:20px;margin-left: 30px"
                 @click="commitMission">
        生成测试用例
      </el-button>
    </div>
    <div>

      <el-col :span="14" style="margin-top: 30px">
        <codemirror :value="codeContent" :options="cmOptions"></codemirror>
      </el-col>
      <p v-if="loading">生成中..</p>
      <p v-else>用例列表：</p>
      <el-col v-loading="loading" :span="8" style="margin-left: 30px;margin-top: 10px" class="case-style">
        <ul class="infinite-list" :v-infinite-scroll="load">
          <li v-for="(acase,indx) in caseContentShow" :key="acase" class="infinite-list-item">
            ({{ indx }})--- {{ acase }}
          </li>
        </ul>
      </el-col>

    </div>
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
      caseContent: [],
      caseContentShow: [],
      caseContentShowLineCount: -3,
      loading: false,
      cmOptions: {
        lineNumbers: true,
        matchBrackets: true,
        readOnly: true,
        mode: 'text/x-c',
        tabSize: 4,
        line: true,
        theme: "base16-light"
      },
      //webscoket
      path: "ws://localhost:8999/websocket/",
      ws: {},


    };
  },
  mounted() {
    this.projectName = this.$route.params.projectName
    this.getSelectMenu()
    //this.initWebsocket()
  },
  computed: {
    noMoreCases() {
      return this.caseContentShow.length + 10 < this.caseContentShowLineCount
    },
    // disabled__() {
    //   console.log("this.noMoreCases:", this.noMoreCases)
    //   return this.noMoreCases
    // }
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
      this.$axios
          .get("/geneCases", {
            params: {
              projectName: this.projectName,
            }
          })
          .then((response) => {
            // console.log(response)

            if (response.data.status === "submit ok") {
              this.$message.success("提交成功")
              this.loading = true
            } else if (response.data.status === "submit error") {
              this.$message.error("提交失败")
            }
          })
          .catch((error) => {
            console.log(error)
          })
      this.timer = window.setInterval(() => {
        setTimeout(() => {
          this.consultMission()
        }, 0)
      }, 5000)


    },
    initWebsocket() {
      // 实例化socket，这里的实例化直接赋值给this.ws是为了后面可以在其它的函数中也能调用websocket方法，例如：this.ws.close(); 完成通信后关闭WebSocket连接
      this.ws = new WebSocket(this.path)

      this.ws.onopen = () => {
        console.log('ws连接状态：' + this.ws.readyState)
        //连接成功则发送一个数据
        this.ws.send('first__连接成功');
      }

      this.ws.onmessage = (data) => {
        console.log('接收到来自服务器的消息：')
        console.log(data)
      }
      this.ws.onclose = () => {
        console.log('ws连接状态：' + this.ws.readyState);
      }

      this.ws.onerror = (error) => {
        console.log(error)
      }
    },
    consultMission() {
      this.$axios
          .get("/consultFileContent", {
            params: {
              projectName: this.projectName,
            }
          })
          .then((response) => {
            console.log(response)

            if (response.data.status === "ok") {
              this.loading = false
              this.caseContent = response.data.cases
              this.caseContentShow = Object.values(this.caseContent).slice(0, this.caseContentShowLineCount)
              window.clearInterval(this.timer)
            } else if (response.data.status === "read file err") {
              this.$message.error("获取失败")
            }
          })
          .catch((error) => {
            console.log(error)
          })
    },
    load() {
      this.caseContentShowLineCount += 3
      this.caseContentShow = Object.values(this.caseContent).slice(0, this.caseContentShowLineCount)
      console.log("3333-==-=", this.caseContent)
      console.log("caseContentShowLineCount:", this.caseContentShowLineCount)
      // console.log("caseContentShowLineCount:", this.caseContentShowLineCount)

    }
  },


}
</script>

<style>
.case-style {
  background-color: #E6CEAC;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  font-size: 20px;
  color: #000000;

}


.CodeMirror {
  height: 620px;
  /*display: inline-block;*/
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  font-size: 20px;
  float: left;
  width: 100%;
}


.infinite-list {
  overflow: auto;
  height: 500px
}

.infinite-list-item {
  font-size: 20px;
}
</style>