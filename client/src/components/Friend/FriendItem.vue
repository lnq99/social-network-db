<template>
  <card class="row hover p18">
    <span>
      <el-avatar class="ava" :size="40" :src="friend.avatars"></el-avatar
    ></span>
    <span>{{ friend.name }}</span></card
  >
</template>

<script>
export default {
  props: ['friend'],
}
</script>

<style scoped>
.row {
  display: flex;
}
.ava {
  margin: -10px 20px -10px -10px;
}
</style>
