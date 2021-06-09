import axios from 'axios'

export default {
  namespaced: true,
  state: {
    notif: [],
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
