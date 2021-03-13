import Vue from 'vue'
import {getToken} from '@/utils/auth'

Vue.prototype.$getAccessToken = function () {
  return getToken()
}

Vue.prototype.$getVueAppBaseAPI = function () {
  return process.env.VUE_APP_BASE_API
}
