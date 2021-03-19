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
    {
      path: '/views/:projectName/caseGenerate',
      name: 'caseGenerate',
      component: () => import('@/components/CaseGenerate')
    },
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
