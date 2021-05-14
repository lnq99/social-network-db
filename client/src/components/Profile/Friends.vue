<template>
  <h2>Friends</h2>
  <div class="card-inner">
    <grid :items="firstFriends(7)">
      <template v-slot="slotProps">
        <el-avatar shape="square" :size="70"></el-avatar>
        <p class="friend-name">Name friend {{ slotProps.item.name }}</p>
      </template>
    </grid>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters({ firstFriends: 'friends/get' }),
  },
}
</script>

<style scoped>
.friend-name {
  margin-top: 0;
  margin-bottom: 0;
  font-size: 0.8em;
  /* white-space: nowrap;
  max-width: 80px;
  overflow: hidden;
  text-overflow: ellipsis; */
}
</style>
