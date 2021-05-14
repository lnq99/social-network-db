<template>
  <div class="ava">
    <div>
      <el-avatar class="avatar" :size="200" :src="avatar"></el-avatar>
    </div>
    <h2 class="name">Username</h2>
  </div>
</template>

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
