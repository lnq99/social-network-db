<template>
  <el-tree
    :data="data"
    :props="defaultProps"
    @node-click="handleNodeClick"
  ></el-tree>
  <comment-input class="cmt-input"></comment-input>
</template>

<script>
import { mapGetters } from 'vuex'
import CommentInput from './CommentInput.vue'

export default {
  components: { CommentInput },
  data() {
    return {
      defaultProps: {
        children: 'children',
        label: 'label',
      },
    }
  },
  computed: {
    ...mapGetters({
      data: 'cmts/cmts',
    }),
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
  padding: 0 18px 5px 18px;
}
.cmt-input {
  padding: 0 18px 10px 18px;
}
</style>
