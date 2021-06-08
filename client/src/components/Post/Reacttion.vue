<template>
  <div class="reacts">
    <div class="react" v-for="(v, i) in reaction">
      {{ v }}
      <img :src="icons[i]" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  props: ['reaction'],
  computed: {
    ...mapState('reaction', ['like', 'love', 'haha', 'wow', 'sad', 'angry']),
  },
  created() {
    this.icons = [
      this.like, this.love, this.haha, this.wow, this.sad, this.angry
    ]
  }
}
</script>
<style scoped>
.reacts {
  display: inline-flex;
  margin-bottom: 0;
  opacity: 0.7;
}
.react {
  display: flex;
  justify-content: flex-end;
  font-size: 0.9em;
  font-weight: lighter;
  width: 40px;
}
img {
  margin-left: 4px;
  width: 1em;
}
</style>
