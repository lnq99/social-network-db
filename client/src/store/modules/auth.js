import axios from 'axios'

export default {
  state: {
    isLoggedIn: false,
    token: '',
  },
  mutations: {
    auth(state, loginStatus) {
      state.isLoggedIn = loginStatus
    },
    setToken(state, token) {
      state.token = token
      // axios.defaults.headers.common['Authorization'] = token
    },
  },
  actions: {
    async login({ commit, state }, payload) {
      let options = {
        method: 'POST',
        url: '/api/login',
        data: payload || '',
      }

      return axios(options).then((data) => {
        if (data) {
          commit('setRootId', data.user.id)
          commit('setToken', data.token)
          commit('auth', true)
          commit('profile/initProfile', data.user, { root: true })
          console.log(data)
          return true
        }
      })
    },
    signup(_, data) {
      data.birthdate = data.birthdate.toISOString()
      data.gender = data.gender[0]
      console.log(data)
      let options = {
        method: 'POST',
        url: '/api/register',
        data: data,
      }
      return axios(options)
    },
    logout({ commit }) {
      axios({ url: '/api/logout' }).then(() => {
        commit('auth', false)
        commit('setToken', '')
      })
    },
  },
  getters: {
    isAuthenticated(state) {
      return state.isLoggedIn
    },
  },
}