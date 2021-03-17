<template>
  <div>
    我是grades页面
    {{ stuInfo }}
  </div>
</template>

<script>

export default {
  name: "Grade",
  data() {
    return {
      projectName:"grade",
      stuInfo: null
    }
  },
  created() {
    console.log("start axios...")
    this.$axios
        .post("/getDirInfo", {
          "projectName":this.projectName
        })
        .then((response) => {
          this.stuInfo = JSON.stringify(response.data)
          console.log(response.data)
        })
        .catch((error) => {
          console.log(error)
        })
  }
}
</script>

<style scoped>

</style>