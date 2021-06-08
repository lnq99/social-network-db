<template>
  <scroll-container :items="friends" class="friend-list">
    <template v-slot="slotProps">
      <friend-item :friend="slotProps.item"></friend-item>
    </template>
  </scroll-container>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import FriendItem from '../Friend/FriendItem.vue'

export default {
  components: { FriendItem },
  computed: {
    ...mapState('profile', ['friends']),
  },
  methods: {
    ...mapActions({ getFriends: 'profile/getFriends' })
  },
  created() {
    this.getFriends()
  }
}
</script>

<style scoped>
.friend-list {
  height: 90vh;
  overflow: auto;
}
</style>
