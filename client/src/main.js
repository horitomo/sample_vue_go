import Vue from 'vue'
import App from './App.vue'

import router from './router'
import store from './store'

require('@/assets/sass/main.scss')

import axios from 'axios' //追記
import VueAxios from 'vue-axios' //追記

Vue.use(VueAxios, axios) //追記

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
