<template>
  <el-tree
    v-if="loaded"
    node-key="id"
    :data="data"
    :props="defaultProps"
    @node-click="handleNodeClick"
  >
    <template #default="{ node, data }">
      <div class="cmt">
        <short-info :id="data.userId">
          <template v-slot="slotProps">
            <span>
              <el-avatar
                class="ava"
                :size="30"
                :src="slotProps.avatars"
              ></el-avatar
            ></span>
            <span>
              <card class="cmt-box p12 box-shadow">
                <div class="cmt-header-r">
                  <span class="cmt-author">{{ slotProps.uname }}</span>
                  <time class="cmt-time">{{ data.created }}</time>
                </div>
                <div>{{ node.label }}</div>
                <!-- <span>
              <a @click="append(data)"> Append </a>
              <a @click="remove(node, data)"> Delete </a>
            </span> -->
              </card></span
            >
          </template>
        </short-info>
      </div>
    </template>
  </el-tree>
  <comment-input class="cmt-input"></comment-input>
</template>

<script>
import { mapActions } from 'vuex'
import CommentInput from './CommentInput.vue'

export default {
  props: ['postId'],
  components: { CommentInput },
  data() {
    return {
      loaded: false,
      defaultProps: {
        children: 'children',
        label: 'content',
      },
    }
  },
  methods: {
    ...mapActions({ getCmtTree: 'cmt/getCmtTree' }),
    handleNodeClick(data) {
      console.log(data)
    },
  },
  created() {
    console.log(this.postId)
    this.getCmtTree(this.postId).then(res => {
      this.data = res.data
      this.loaded = true
    })
  },
}
</script>

<style scoped>
.el-tree {
  background: none;
  padding: 0 18px 0 18px;
}
.el-tree-node {
  margin: 5px;
}
.cmt-box {
  /* margin-left: 46px; */
}
.cmt {
  display: flex;
  padding-bottom: 4px;
  opacity: 0.9;
}
.cmt-header-r {
  font-size: 0.9em;
  margin-bottom: 4px;
}
.cmt-author {
  font-weight: 600;
}
.cmt-time {
  padding-left: 12px;
  opacity: 0.8;
  font-size: 0.8em;
}
.ava {
  margin-right: 8px;
}
.card {
  display: inline-block;
}
</style>
