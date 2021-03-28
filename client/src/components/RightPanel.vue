<template>
  <el-aside width="20vw">
    <el-row>
      <p>
        For a guide and recipes on how to configure / customize this project,<br />
        check out the
        <a href="https://cli.vuejs.org" target="_blank" rel="noopener"
          >vue-cli documentation</a
        >.
      </p>
    </el-row>
    <hr />
    <el-row>
      <friend-list></friend-list>
    </el-row>
  </el-aside>
</template>

<script>
import FriendList from './FriendList.vue'
export default {
  components: { FriendList },
  data() {
    const item = {
      date: '2016-05-02',
      name: 'Tom',
    }
    return {
      tableData: Array(20).fill(item),
    }
  },
}
</script>

<style>
.el-aside {
  width: 20vw;
  padding: 10px;
}
</style>
