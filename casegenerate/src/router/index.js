import Vue from 'vue'
import Router from 'vue-router'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: () => import('@/components/HelloWorld')
    },
    {
      path: '/views/:projectName',
      name: 'public',
      component: () => import('@/components/Public')
    },
    // {
    //   path: '/views/digits',
    //   name: 'digits',
    //   component: () => import('@/components/Digits')
    // },
    // {
    //   path: '/views/median',
    //   name: 'median',
    //   component: () => import('@/components/Median')
    // },{
    //   path: '/views/smallest',
    //   name: 'smallest',
    //   component: () => import('@/components/Smallest')
    // },{
    //   path: '/views/syllables',
    //   name: 'syllables',
    //   component: () => import('@/components/Syllables')
    // },{
    //   path: '/views/checksum',
    //   name: 'checksum',
    //   component: () => import('@/components/Checksum')
    // }
  ]
})
