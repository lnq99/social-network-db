<template>
  <card class="row hover p18">
    <el-avatar class="ava" :size="40" :src="friend.ava"></el-avatar>
    <span>{{ friend.name }}</span></card
  >
</template>

<script>
export default {
  props: ['friend'],
}
</script>

<style scoped>
.ava {
  margin: -10px 20px -10px -10px;
}
</style>
