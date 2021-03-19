<template>
  <div id="app">
    <h2 class="h2-class">
      <img :src="logoPath" alt="" style="width: 60px;height: 60px">
      introClass
    </h2>
    <el-menu
        mode="horizontal"
        background-color="#545c64"
        text-color="#ffffff"
        active-text-color="#ffd04b"
        router
        :default-active="this.$route.path">
      <el-menu-item v-for="(item,i) in itemList" :key="i" :index="item.copUrl">
        <template slot="title">
          <i class="el-icon-tickets"></i>
          <span class="el-calendar__title"> {{ item.copName }}</span>
        </template>
      </el-menu-item>
    </el-menu>
    <router-view class="menu-right"/>
  </div>

</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'  //导入组件
// import Digits from '@/components/Digits.vue'
// import Grades from '@/components/Grade.vue'
export default {
  name: 'App',
  data() {
    return {
      logoPath:require('@/statics/icon/logo.png'),
      itemList: [
        {copUrl: '/views/grade', copName: 'grade'},
        {copUrl: '/views/digits', copName: 'digits'},
        {copUrl: '/views/median', copName: 'median'},
        {copUrl: '/views/smallest', copName: 'smallest'},
        {copUrl: '/views/syllables', copName: 'syllables'},
        {copUrl: '/views/checksum', copName: 'checksum'}
      ]
    };
  },

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.el-calendar__title {
  font-size: 25px;
}
.h2-class{
  margin-top: 15px;
  font-size: 50px;
  margin-bottom: 5px;
}
.el-menu{
  margin-top: 10px;
}

</style>
