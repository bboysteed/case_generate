<template>
  <div>
    我是{{this.projectName}}页面
    <ul>
      <li v-for="(stu,index) in stuInfo" :key="index">
        {{ stu.stu_name | enShort(index)}}
        <ul>
          <li v-for="(job,ind2) in stu.jobs_names" :key="ind2">
            {{job}}
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
      projectName: null ,
      stuInfo: null
    }
  },
  mounted() {
    this.projectName=this.$route.params.projectName
    console.log(this.projectName)
    this.fetchData()
  },
  watch:{
    $route(){
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
            // console.log(this.stuInfo)
          })
          .catch((error) => {
            console.log(error)
          })
    }

  },
  filters:{
    enShort(value,index) {
      return "student-".concat(index).concat("--(",value.toString().slice(0,16),")")
    }
  }
}
</script>

<style scoped>

</style>