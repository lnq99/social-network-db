import { createStore } from 'vuex'
import newsfeed from './modules/newsfeed.js'
import notif from './modules/notif.js'
import cmts from './modules/comments.js'
import friends from './modules/friends'
import photos from './modules/photos.js'
import axios from 'axios'

const store = createStore({
  state: {
    isLoggedIn: false,
    isDark: false,
    token: '',
    id: 0,
    name: 'Name',
    avatarL: require('../assets/logo.png'),
    intro: '',
  },
  mutations: {
    auth(state, loginStatus) {
      state.isLoggedIn = loginStatus
    },
    switchTheme(state) {
      state.isDark = !state.isDark
    },
    saveIntro(state, intro) {
      state.intro = intro
    },
    saveToken(state, token) {
      state.token = token
    },
    initProfile(state, profile) {
      state.id = profile.id
      state.name = profile.name
      state.avatarL = profile.avatarL
      state.intro = profile.intro
    },
  },
  actions: {
    async login({ commit, getters }, payload) {
      let options = {
        method: 'POST',
        url: 'api/login',
        headers: getters.header,
        data: {
          email: payload.email,
          password: payload.password,
        },
      }
      let response = await axios(options).catch(() => {})
      let responseOK =
        response && response.status === 200 && response.statusText === 'OK'
      if (responseOK) {
        let data = await response.data
        console.log(data)
        if (data) {
          commit('saveToken', data.token)
          commit('initProfile', data.user)
          commit('auth', true)
          return true
        }
      }

      return false
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
    saveIntro({ commit }, intro) {
      commit('saveIntro', intro)
    },
  },
  getters: {
    isAuthenticated(state) {
      return state.isLoggedIn
    },
    intro(state) {
      return state.intro
    },
    header(state) {
      return {
        Accept: 'application/json',
        'Content-Type': 'application/json;charset=UTF-8',
        Authorization: `Bearer ${state.token}`,
      }
    },
  },
  modules: { newsfeed, notif, cmts, friends, photos },
})

export default store
