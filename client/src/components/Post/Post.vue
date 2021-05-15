<template>
  <card class="post card-hl">
    <div class="post-header">
      <el-avatar :size="40" :src="avatar"></el-avatar>
      <div class="post-header-r">
        <span class="post-author">{{ name }}</span>
        <time class="post-time">{{ new Date() }}</time>
      </div>
    </div>
    <p class="post-content">{{ text }}</p>
    <div class="attach">
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

<script>
import Comments from '../Post/Comments.vue'

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

<style lang="scss" scoped>
@import '@/style.scss';

.post {
  margin-bottom: 12px;
}
.post-header {
  display: flex;
  align-items: center;
  text-align: left;
  padding: $p4;
}
.post-header-r {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding-left: $p4;
}
.post-author {
  font-weight: 900;
}
.post-time {
  font-size: 0.8em;
  font-weight: 100;
}

.post-content {
  margin: 0 $p4 $p4 $p4;
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
  padding: 0px $p4 10px $p4;
}

.btn {
  color: var(--fg);
  font-weight: 300;
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
