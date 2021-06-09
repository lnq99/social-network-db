import axios from 'axios'

export default {
  namespaced: true,
  state: {
    like: require('@/assets/like.svg'),
    love: require('@/assets/love.svg'),
    haha: require('@/assets/haha.svg'),
    wow: require('@/assets/wow.svg'),
    sad: require('@/assets/sad.svg'),
    angry: require('@/assets/angry.svg'),
  },
  actions: {
    async getReaction(_, postId) {
      return axios({ url: `/api/react/${postId}` })
        .catch(() => {})
        .then((r) => r.data)
    },
    async getReactionType(_, postId) {
      return axios({ url: `/api/react/u/${postId}` })
        .catch(() => {})
        .then((r) => r.data)
    },
    async react(_, { postId, type }) {
      return axios({ method: 'put', url: `/api/react/${postId}/${type}` })
        .catch(() => {})
        .then((r) => r.data)
    },
  },
}
