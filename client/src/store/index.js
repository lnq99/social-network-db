import { createStore } from 'vuex'
import notif from './modules/notif.js'
import cmt from './modules/cmt.js'
import reaction from './modules/reaction.js'
import relationship from './modules/relationship.js'
import photo from './modules/photo.js'
import profile from './modules/profile.js'
import post from './modules/post.js'
import axios from 'axios'

const store = createStore({
  state: {
    isLoggedIn: false,
    isDark: false,
    token: '',
    id: 0,
    // name: 'Name',
    // avatarl: require('../assets/logo.png'),
    // intro: '',
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
      state.avatarl = profile.avatarl
      state.avatars = profile.avatars
      state.intro = profile.intro
    },
  },
  actions: {
    async login({ commit, getters }, payload) {
      let data
      if (payload) {
        data = {
          email: payload.email,
          password: payload.password,
        }
      }
      let options = {
        method: 'POST',
        url: '/api/login',
        headers: getters.header,
        data,
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
      axios({ url: '/api/logout' })
        .catch((r) => console.log(r))
        .then(() => {
          commit('auth', false)
          commit('saveToken', '')
        })
      // document.cookie.split(';').forEach((c) => {
      //   document.cookie = c
      //     .replace(/^ +/, '')
      //     .replace(/=.*/, '=;expires=' + new Date().toUTCString() + ';path=/')
      //   console.log(c)
      // })
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
  modules: { profile, post, notif, cmt, reaction, relationship, photo },
})

export default store
