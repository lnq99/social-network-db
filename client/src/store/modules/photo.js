import axios from 'axios'

export default {
  namespaced: true,
  state: {},
  actions: {
    async getPhoto(_, id) {
      let options = {
        method: 'GET',
        url: `/api/photo/${id}`,
      }
      return axios(options)
    },
    async getPhotosOfProfile(_, profileId) {
      let options = {
        method: 'GET',
        url: `/api/photo/u/${profileId}`,
      }
      return axios(options)
    },
  },
}
