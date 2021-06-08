import axios from 'axios'

export default {
  namespaced: true,
  state: {
    feed: [],
    friends: [],
    shortProfiles: {},
  },
  // state: {
  //   profile: {
  //     id: '605a446c0a3330c15c1ee9a7',
  //     index: 0,
  //     avatar: 'http://placehold.it/32x32',
  //     photo: 'http://placehold.it/1280x720',
  //     age: 24,
  //     name: 'Kristy Mccoy',
  //     text:
  //       'Pariatur cupidatat magna et in elit sit tempor occaecat anim qui dolor. Aliqua velit qui nisi in cillum do adipisicing velit cupidatat proident. Reprehenderit eiusmod ipsum proident irure laboris anim. Enim laborum amet sit mollit minim ea. Excepteur excepteur occaecat est labore cillum enim aliqua excepteur ad incididunt.\r\n',
  //     time: '2015-02-07T10:13:05 -03:00',
  //     tags: ['aute', 'tag'],
  //   },
  // },
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
    async getFriends({ state, commit }) {
      return axios({ url: '/api/friends' })
        .catch(() => {})
        .then((res) => {
          state.friends = res.data
          commit('cacheShortProfileArray', res.data)
        })
    },
  },
}
