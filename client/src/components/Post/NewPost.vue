<template>
  <card class="new-post card-hl">
    <div class="new-post-content">
      <el-avatar :size="40" :src="avatar"></el-avatar>
      <div class="input">
        <span @click="createPost = true">
          Create a post
        </span>
      </div>
    </div>
  </card>
  <el-dialog title="Create Post" v-model="createPost" width="40%" center>
    <el-input
      type="textarea"
      :autosize="{ minRows: 2, maxRows: 4 }"
      placeholder="Please input"
      v-model="postContent"
    >
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="createPost = false">Post</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import { ref } from 'vue'
export default {
  data() {
    return {
      createPost: false,
      postContent: ref(''),
    }
  },
}
</script>

<style lang="scss" scoped>
@import '@/style.scss';

.new-post {
  width: 100%;
  margin-bottom: 10px;
}
.new-post-content {
  padding: $p4;
  display: flex;
  align-content: center;
}
.input {
  padding-left: $p4;
  width: auto;
  background: var(--hl);
  border: none;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>
