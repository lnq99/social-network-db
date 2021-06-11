import axios from 'axios'

export default {
  namespaced: true,
  actions: {
    async getCmtTree(_, postId) {
      return axios({ url: `/api/cmt/${postId}` })
    },
    async comment(_, cmtBody) {
      let options = {
        method: 'POST',
        url: `/api/cmt`,
        data: cmtBody,
      }
      return axios(options)
    },
  },
}
