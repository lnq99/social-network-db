<template>
  <div class="ava">
    <div>
      <el-avatar class="avatar" :size="200" :src="avatar"></el-avatar>
    </div>
    <h2 class="name">{{ name }}</h2>
  </div>
</template>

<script>

export default {
  props: ['name', 'avatar'],
}
</script>

<style scoped>
.avatar {
  margin: 20px;
}
.name {
  text-align: center;
}
.ava {
  display: flex;
  flex-direction: column;
  text-align: center;
}
</style>
