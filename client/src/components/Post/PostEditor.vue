<template>
  <el-input
    type="textarea"
    :autosize="{ minRows: 6, maxRows: 10 }"
    placeholder="Please input"
    v-model="postContent"
  >
  </el-input>

  <el-upload
    class="upload center"
    drag
    :on-success="handleSuccess"
    :on-preview="handlePreview"
    :on-remove="handleRemove"
    :file-list="fileList"
    list-type="picture"
    limit="1"
    :disabled="!show"
  >
    <i class="el-icon-upload"></i>
    <div class="el-upload__text">
      Drop file here or <em>click to upload</em>
    </div>
    <template #tip>
      <div class="el-upload__tip">
        jpg/png files with a size less than 500kb
      </div>
    </template>
  </el-upload>
  <el-button type="primary" @click="onCreatePost = false">Post</el-button>
</template>

<script>
import { ref } from "vue"
export default {
  data() {
    return {
      show: true,
      postContent: ref(""),
      fileList: [],
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList)
      this.show = true
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    handleSuccess(file, fileList) {
      console.log(file, fileList)
      this.show = false
    },
  },
};
</script>

<style scoped>
.upload {
  margin: 20px;
}
</style>
