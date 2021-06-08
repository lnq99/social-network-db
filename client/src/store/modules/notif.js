import axios from 'axios'

export default {
  namespaced: true,
  state: {
    notif: [
      // { text: 'Join like you post1', time: '2m', link: 'http://ok' },
      // { text: 'Join like you post2', time: '2m', link: 'http://ok' },
    ],
  },
  actions: {
    getNotif({ state }) {
      return axios({ url: '/api/notif' })
        .catch(() => {})
        .then((r) => {
          state.notif = r.data
        })
    },
  },
  getters: {
    notif(state) {
      return state.notif
    },
  },
}
