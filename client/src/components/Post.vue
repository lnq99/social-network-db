<template>
  <card class="post" :body-style="{ padding: '0px' }" shadow="never">
    <div class="post-header">
      <el-avatar class="avatar-s" :size="40" :src="avatar"></el-avatar>
      <div class="post-header-r">
        <span class="post-author">{{ name }}</span>
        <time class="post-time">{{ new Date() }}</time>
      </div>
    </div>
    <p class="post-content">{{ text }}</p>
    <div class="attack">
      <img :src="photo" />
    </div>
    <hr />
    <el-row class="row">
      <el-button
        class="btn"
        :class="{ 'btn-on': liked }"
        type="primary"
        @click="onLike"
        size="small"
        >Like</el-button
      >
      <el-button class="btn" size="small" @click="onComment">Comment</el-button>
    </el-row>
    <comments v-if="comment"></comments>
  </card>
</template>
// :type="liked ? 'primary' : 'default'"
<script>
import Comments from './Comments.vue'

export default {
  components: { Comments },
  props: ['id', 'avatar', 'name', 'text', 'time', 'photo', 'tag'],
  data() {
    return {
      liked: false,
      comment: false,
    }
  },
  methods: {
    onLike() {
      this.liked = !this.liked
    },
    onComment() {
      this.comment = !this.comment
    },
  },
}
</script>

<style scoped>
.post {
  margin-top: 10px;
  background: #161b22;
  border-radius: 10px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}
.post-header {
  display: flex;
  align-items: center;
  text-align: left;
  padding: 18px;
}
.post-header-r {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding-left: 18px;
}
.post-author {
  font-weight: 900;
}
.post-time {
  font-size: 0.8em;
  font-weight: 100;
}

.post-content {
  margin: 0 18px 18px 18px;
  text-align: left;
}
img {
  width: 100%;
  display: block;
}

.row {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 0px 18px 10px 18px;
}

.btn {
  color: #c9d1d9;
  width: max(25%, 100px);
}

.btn,
.btn:focus {
  background: none;
}

.btn-on,
.btn-on:focus {
  background: #409eff;
}
</style>
