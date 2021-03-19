<template>
  <div>
    <el-container>
      <el-aside width="200px">
        {{ this.projectName }}
      </el-aside>
      <el-container>
        <el-header>
          功能列表
        </el-header>
        <el-main>

          <el-button class="btn" icon="el-icon-setting" round type="primary"
                     @click="$router.push('/views/'+projectName+'/caseGenerate')">用例生成
          </el-button>
          <el-button class="btn" icon="el-icon-edit-outline" round type="success"
                     @click="$router.push('/views/'+projectName+'/whiteBox')">白盒测试
          </el-button>
          <el-button class="btn" icon="el-icon-document" round type="info"
                     @click="$router.push('/views/'+projectName+'/whiteBox')">报告分析
          </el-button>


        </el-main>
      </el-container>

    </el-container>

    <ul>
      <li v-for="(stu,index) in stuInfo" :key="index">
        {{stu.label}}
        <ul>
          <li v-for="(job,ind1) in stu.children" :key="ind1">
            {{job.label}}
          </li>
        </ul>
      </li>

    </ul>

  </div>
</template>

<script>

export default {
  name: "Public",
  data() {
    return {
      projectName: null,
      stuInfo: null
    }
  },
  mounted() {
    this.projectName = this.$route.params.projectName
    console.log(this.projectName)
    this.fetchData()
  },
  watch: {
    $route() {
      this.projectName = this.$route.params.projectName
      this.fetchData()
    }
  },
  methods: {
    fetchData() {
      console.log("start axios...")
      this.$axios
          .post("/getDirInfo", this.$qs.stringify(
              {
                "projectName": this.projectName
              }
          ))
          .then((response) => {
            this.stuInfo = response.data.message.all_students
            console.log(this.stuInfo)
          })
          .catch((error) => {
            console.log(error)
          })
    }

  },
  filters: {
    enShort(value, index) {
      return "student-".concat(index).concat("--(", value.toString().slice(0, 16), ")")
    }
  }
}
</script>

<style scoped>
.btn {
  width: 270px;
  height: 120px;
  font-size: 30px;
  /*margin: 0 auto;*/
  display: inline-block;
  /*float: left;*/
  margin-left: 97px;

}


.el-header, .el-footer {
  background-color: #dddddd;
  color: #000000;
  text-align: center;
  line-height: 60px;
  font-size: 30px;
  font-weight: bold;
}

.el-aside {
  background-color: #528B8B;
  color: #FFD700;
  text-align: center;
  line-height: 200px;
  font-size: 30px;
}

/*.el-main {*/
/*  background-color: #E9EEF3;*/
/*  color: #333;*/
/*  text-align: center;*/
/*  line-height: 160px;*/
/*}*/

body > .el-container {
  margin-bottom: 40px;
  margin-top: 20px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>