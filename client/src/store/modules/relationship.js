import axios from 'axios'

export default {
  namespaced: true,
  state: {
    friends: [],
  },
  actions: {
    async getMutualFriends(_, id) {
      return axios({ url: `/api/rel/mutual-friends?id=${id}` })
        .catch(() => {})
        .then((res) => res.data)
    },
    async getFriends({ rootState, state, commit }, id) {
      if (id == rootState.id && state.friends.length > 0) {
        return state.friends
      }
      return axios({ url: `/api/rel/friends/${id}` })
        .catch(() => {})
        .then((res) => {
          if (id == rootState.id && state.friends.length == 0)
            state.friends = res.data
          commit('profile/cacheShortProfileArray', res.data, { root: true })
          return res.data
        })
    },
  },
  getters: {
    // friends(state) {
    //   let a = [...new Array(20)]
    //     .map((i) => Object.create(state.friend))
    //     .map((e, i) => {
    //       e.name = `${e.name} ${i}`
    //       return e
    //     })
    //   console.log(state.friend)
    //   return a
    // },
    // get(_, getters) {
    //   return (n) => getters.friends.slice(0, n)
    // },
  },
}
