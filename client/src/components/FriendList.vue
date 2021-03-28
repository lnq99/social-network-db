<template>
  <div class="friend-list">
    <friend-item v-for="friend in friends" :name="friend.name"></friend-item>
  </div>
</template>

<script>
import FriendItem from './FriendItem.vue'

export default {
  components: { FriendItem },
  data() {
    const item = {
      date: '2016-05-02',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
    }
    return {
      friends: Array(20).fill(item),
    }
  },
}
</script>

<style scoped>
.friend-list {
  width: 100%;
}
</style>
