<template>
  <h2>Photos <a class="btn-right" href="/photo">See more</a></h2>
  <grid :items="firstPhotos(7)">
    <template v-slot="slotProps">
      <img class="preview" :src="slotProps.item" />
    </template>
  </grid>
</template>

<script>
import { mapGetters } from 'vuex'
import { squarePreview } from '@/utils.js'

export default {
  computed: {
    ...mapGetters({ firstPhotos: 'photos/get' }),
  },
  mounted() {
    squarePreview()
  },
}
</script>

<style scoped>
.flex-container {
  display: flex;
  flex-flow: row wrap;
}
.btn-right {
  float: right;
  font-weight: lighter;
  font-size: 0.8em;
  margin-top: 0.1em;
}
</style>
