<template>
  <card class="new-post card-hl">
    <div class="new-post-content">
      <el-avatar :size="40" :src="avatar"></el-avatar>
      <div class="input">
        <span @click="onCreatePost = true"> Create a post </span>
      </div>
    </div>
  </card>
  <el-dialog
    title="Create Post"
    v-model="onCreatePost"
    width="max(40%, 400px)"
    center
  >
    <post-editor></post-editor>
  </el-dialog>
</template>

<script>
import PostEditor from "./PostEditor.vue"
export default {
  components: { PostEditor },
  data() {
    return {
      onCreatePost: false,
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/style.scss";

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
  margin-left: $p4;
  background: var(--hl);
  border: none;
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 1;
}
.input:hover {
  background: var(--hl2);
}
</style>
