import axios from 'axios'

export default {
  namespaced: true,
  state: {
    photos: [
      'http://placehold.it/1080x720',
      'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg',
      'http://placehold.it/1080x720',
      'https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg',
    ],
  },
  actions: {
    async getPhoto(_, id) {
      let options = {
        method: 'GET',
        url: `/api/photo/${id}`,
      }
      return axios(options)
        .catch(() => {})
        .then((r) => r.data)
    },
    async getPhotosOfProfile(_, profileId) {
      let options = {
        method: 'GET',
        url: `/api/photo/u/${profileId}`,
      }
      return axios(options)
        .catch(() => {})
        .then((r) => r.data)
    },
  },
  getters: {
    photos(state) {
      return state.photos
    },
    get(state) {
      return (n) => {
        if (n) return state.photos.slice(0, n)
        else return state.photos
      }
    },
  },
}
