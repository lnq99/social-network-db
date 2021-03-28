import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

import Card from '@/components/Base/Card.vue'
import LinkCard from '@/components/Base/LinkCard.vue'

// Vue.config.productionTip = false

const app = createApp(App)

app.use(router)
app.use(store)

app.component('Card', Card)
app.component('LinkCard', LinkCard)

app.use(ElementPlus)

app.mount('#app')
