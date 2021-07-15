import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './components/index'
import i18n from './i18n'
import mixin from './mixin'
import 'normalize.css'
import './style/index.scss'
Vue.config.productionTip = false
Vue.mixin(mixin)
new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')
