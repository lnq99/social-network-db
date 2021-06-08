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
  getters: {
    posts(state) {
      let posts = Array(10).fill(state.post)
      return posts
    },
  },
}
