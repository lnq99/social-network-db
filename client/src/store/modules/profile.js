import axios from 'axios'

export default {
  namespaced: true,
  state: {
    feed: [],
    shortProfiles: {},
  },
  mutations: {
    cacheShortProfileArray(state, arr) {
      for (let p of arr) {
        state.shortProfiles[p.id] = p
      }
    },
    cacheShortProfile(state, p) {
      state.shortProfiles[p.id] = p
    },
  },
  actions: {
    async getProfile(_, id) {
      let options = {
        method: 'GET',
        baseURL: '',
        url: `/api/profile/${id}`,
      }
      return axios(options)
        .catch((err) => {
          console.log(err)
        })
        .then((r) => r.data)
    },
    async getProfileShort({ state, commit }, id) {
      let p = state.shortProfiles[id]
      if (p) {
        // console.log('cache hit', id)
        return p
      }
      let options = {
        method: 'GET',
        url: `/api/profile/short/${id}`,
      }
      return axios(options)
        .catch(() => {})
        .then((res) => {
          // console.log('cache missed', id)
          commit('cacheShortProfile', res.data)
          return res.data
        })
    },
    async searchProfile(_, key) {
      let options = {
        method: 'GET',
        url: `/api/search?k=${key}`,
      }
      return axios(options)
        .catch(() => {})
        .then((res) => {
          return res.data
        })
    },
  },
}
