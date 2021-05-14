export default {
  namespaced: true,
  state: {
    friend: {
      date: '2016-05-02',
      name: 'Friend',
      address: 'No. 189, Grove St, Los Angeles',
      ava: 'https://via.placeholder.com/40',
    },
  },
  getters: {
    friends(state) {
      let a = [...new Array(20)]
        .map((i) => Object.create(state.friend))
        .map((e, i) => {
          e.name = `${e.name} ${i}`
          return e
        })
      console.log(state.friend)
      return a
    },
    get(_, getters) {
      return (n) => getters.friends.slice(0, n)
    },
  },
}
