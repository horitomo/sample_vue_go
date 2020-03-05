import Vue from 'vue'
import App from './App.vue'

import router from './router'
import Vuex from 'vuex'

require('@/assets/sass/main.scss')

import axios from 'axios' //追記
import VueAxios from 'vue-axios' //追記

Vue.use(VueAxios, axios, Vuex) //追記

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
