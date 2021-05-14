<template>
  <el-row>
    <search></search>
  </el-row>
  <hr />
  <friend-list></friend-list>
</template>

<script>
import FriendList from './Friend/FriendList.vue'
import Search from './Search/Search.vue'

export default {
  components: { FriendList, Search },
}
</script>
