import axios from 'axios'

export default {
  namespaced: true,
  actions: {
    async getCmtTree(_, postId) {
      let options = {
        method: 'GET',
        url: `/api/cmt/${postId}`,
      }
      return axios(options).catch(() => {
        console.log('err')
      })
    },
    async comment(_, postBody) {
      let options = {
        method: 'POST',
        url: `/api/cmt`,
        data: postBody,
      }
      return axios(options).catch((err) => {
        console.log(err)
      })
    },
  },
}
