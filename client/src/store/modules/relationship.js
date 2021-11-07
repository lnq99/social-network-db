import axios from 'axios'

export default {
  namespaced: true,
  state: {
    friends: [],
  },
  actions: {
    // async getMutualFriends(_, id) {
    //   console.log('--', id)
    //   return axios({ url: `/api/rel/mutual-friends?id=${id}` })
    // },
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
    async changeRelationship(_, [userId, type]) {
      return axios({
        method: 'PUT',
        url: `/api/rel/${userId}/${type}`,
      })
    },
    async getMutualAndType(_, id) {
      console.log(id)
      return axios({ url: `/api/rel/mutual-type/${id}` })
    },
  },
}
