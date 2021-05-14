<template>
  <div class="commentinput">
    <el-avatar class="ava" :size="30" :src="avatar"></el-avatar>
    <!-- <div class="input">
      <span @click="createPost = true">
        Create a post
      </span>
    </div> -->
    <el-input
      class="input"
      type="textarea"
      :autosize="{ minRows: 1, maxRows: 4 }"
      placeholder="Write a comment..."
      v-model="cmtModel"
    ></el-input>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  data() {
    return {
      cmtModel: ref(''),
    }
  },
}
</script>
<style scoped>
.commentinput {
  display: flex;
  align-content: center;
  justify-content: center;
}

.ava {
  float: left;
}

.input {
  margin-left: 10px;
  width: 90%;
  /* background: var(--hl); */
}
.input .el-textarea__inner {
  background-color: #888;
}
</style>
