import { createStore } from 'vuex'
import newsfeed from './modules/newsfeed.js'
import notif from './modules/notif.js'
import cmts from './modules/comments.js'
import friends from './modules/friends'
import photos from './modules/photos.js'

const store = createStore({
  state: {
    isLoggedIn: false,
    isDark: false,
    avatarL: require('../assets/logo.png'),
    name: 'User Name',
    info: `Works at ...
    Worked at ...
    Studied at Bauman Moscow State Technical University
    Went to ...
    Lives in Hue, Vietnam
    From Hue, Vietnam
    Joined February 2015`,
  },
  mutations: {
    auth(state, loginStatus) {
      state.isLoggedIn = loginStatus
    },
    switchTheme(state) {
      state.isDark = !state.isDark
    },
    saveInfo(state, info) {
      state.info = info
    },
  },
  actions: {
    async login({ commit }, payload) {
      console.log(payload)
      let validCredentials = {
        email: 'admin@gmail.com',
        password: '12345678',
      }

      await new Promise((resolve) => {
        setTimeout(resolve, 800)
      })

      if (
        payload.email === validCredentials.email &&
        payload.password === validCredentials.password
      ) {
        commit('auth', true)
        return true
      } else {
        return false
      }
    },
    signup(_, payload) {
      console.log(payload)
    },
    logout({ commit }) {
      commit('auth', false)
    },
    switchTheme({ commit }) {
      commit('switchTheme')
    },
    saveInfo({ commit }, info) {
      console.log(info)
      commit('saveInfo', info)
    },
  },
  getters: {
    isAuthenticated(state) {
      return state.isLoggedIn
    },
    info(state) {
      return state.info
    },
  },
  modules: { newsfeed, notif, cmts, friends, photos },
})

export default store
