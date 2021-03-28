<template>
  <card class="row hover p18">
    <el-avatar
      class="ava"
      :size="40"
      src="https://via.placeholder.com/40"
    ></el-avatar>
    <span>{{ name }}</span></card
  >
</template>

<script>
export default {
  props: ['name'],
  created() {
    console.log(this.name)
  },
}
</script>

<style scoped>
.ava {
  margin: -10px 20px -10px -10px;
}
</style>
