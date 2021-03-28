<template>
  <el-aside width="20vw">
    <el-avatar class="avatar" :size="200" :src="avatar"></el-avatar>
    <h1 class="name">NAME</h1>
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div> -->

    <link-card link="/profile"> Profile </link-card>
    <link-card link="/photo"> Photo </link-card>
    <link-card link="/logout"> Logout </link-card>
  </el-aside>
</template>

<script>
export default {
  data() {
    const item = {
      date: '2016-05-02',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
    }
    return {
      // avatar: require('../assets/logo.png'),
      avatar: 'https://i.pravatar.cc/100',
      tableData: Array(20).fill(item),
    }
  },
}
</script>

<style scoped>
.el-aside {
  width: 20vw;
  padding: 10px;
  box-shadow: 0 0 6px rgba(0, 0, 0, 0.2);
}

.avatar {
  margin: 20px;
}
.name {
  padding-bottom: 40px;
}
</style>
