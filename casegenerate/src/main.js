import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from './router'
import Axios from 'axios'
import qs from 'qs';
import codemirror from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'
Vue.use(codemirror)
Vue.prototype.$qs = qs;
Vue.config.productionTip = false
Vue.use(ElementUI, {size:'small', zIndex: 3000});
Vue.prototype.$axios = Axios
Axios.defaults.baseURL = '/api'

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
