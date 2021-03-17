<template>
  <div>
    我是grades页面
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
  name: "Grade",
  data() {
    return {
      projectName: "grade",
      stuInfo: null
    }
  },
  mounted() {
    this.fetchData()
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