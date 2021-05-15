<template>
  <scroll-container :items="friends" class="friend-list">
    <template v-slot="slotProps">
      <friend-item :friend="slotProps.item"></friend-item>
    </template>
  </scroll-container>
</template>

<script>
import { mapGetters } from 'vuex'
import FriendItem from '../Friend/FriendItem.vue'

export default {
  components: { FriendItem },
  computed: {
    ...mapGetters({ friends: 'friends/friends' }),
  },
}
</script>

<style scoped>
.friend-list {
  height: 90vh;
  overflow: auto;
}
</style>
