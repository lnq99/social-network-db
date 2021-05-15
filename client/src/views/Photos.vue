<template>
  <h1>Photos</h1>
  <grid :items="firstPhotos(7)">
    <template v-slot="slotProps">
      <el-image
        class="preview"
        style="height: 100%"
        fit="cover"
        :src="slotProps.item"
        :preview-src-list="[slotProps.item]"
        lazy
      >
      </el-image>
    </template>
  </grid>
</template>

<script>
import { mapGetters } from 'vuex'
import { squarePreviewDiv } from '@/utils.js'

export default {
  computed: {
    ...mapGetters({ firstPhotos: 'photos/get' }),
  },
  mounted() {
    squarePreviewDiv()
  },
}
</script>
