export default {
  namespaced: true,
  state: {
    notif: [
      { text: 'Join like you post1', time: '2m', link: 'http://ok' },
      { text: 'Join like you post2', time: '2m', link: 'http://ok' },
    ],
  },
  getters: {
    notif(state) {
      return state.notif
    },
  },
}
