import axios from 'axios'

export default {
  namespaced: true,
  state: {
    friends: [],
  },
  actions: {
    async getMutualFriends(_, id) {
      return axios({ url: `/api/rel/mutual-friends?id=${id}` })
    },
    async getFriends({ rootState, state, commit }, id) {
      if (id == rootState.id && state.friends.length > 0) {
        return state.friends
      }
      return axios({ url: `/api/rel/friends/${id}` }).then((data) => {
        if (id == rootState.id && state.friends.length == 0)
          state.friends = data
        commit('profile/cacheShortProfileArray', data, { root: true })
        return data
      })
    },
  },
}
