<template>
  <el-tree
    :data="data"
    :props="defaultProps"
    @node-click="handleNodeClick"
  ></el-tree>
</template>

<script>
import comments from './comments.js'
export default {
  data() {
    return {
      data: comments,
      defaultProps: {
        children: 'children',
        label: 'label',
      },
    }
  },
  methods: {
    handleNodeClick(data) {
      console.log(data)
    },
  },
}
</script>

<style scoped>
.el-tree {
  background: none;
}
</style>
