<template>
  <card v-if="loaded" class="post card-hl">
    <div class="post-header">
      <short-info :id="data.userId">
        <template v-slot="slotProps">
          <el-avatar
            class="ava"
            :size="40"
            :src="slotProps.avatars"
          ></el-avatar>
          <div class="post-header-r">
            <span class="post-author">{{ slotProps.uname }}</span>
            <time class="post-time">{{ data.created }}</time>
          </div>
        </template>
      </short-info>
    </div>
    <p class="post-content">{{ data.content }}</p>
    <div v-if="(data.atchType = 'photo')" class="attach">
      <img :src="data.atchUrl" />
    </div>
    <hr />
    <div class="row">
      <div class="btn-contaner">
        <el-button
          class="btn"
          :class="{ 'btn-on': liked }"
          type="primary"
          @click="onLike"
          size="small"
          >Like</el-button
        >
        <el-button class="btn" size="small" @click="onComment">
          {{ data.cmtCount }} Comments</el-button
        >
      </div>
      <react-cmt
        :reaction="data.reaction"
        :cmtCount="data.cmtCount"
      ></react-cmt>
    </div>
    <div v-if="comment">
      <hr />
      <comments :postId="id"></comments>
    </div>
  </card>
</template>

<script>
import { mapActions } from 'vuex'
import ShortInfo from '../Profile/ShortInfo.vue'
import ReactCmt from './Reacttion.vue'
import Comments from '../Post/Comments.vue'

export default {
  components: { ShortInfo, ReactCmt, Comments },
  props: ['id'],
  data() {
    return {
      data: {},
      author: {},
      liked: false,
      comment: false,
      loaded: false,
    }
  },
  computed: {
  },
  methods: {
    ...mapActions({ getPost: 'post/getPost', getPhoto: 'photo/getPhoto', getProfileShort: 'profile/getProfileShort' }),

    onLike() {
      this.liked = !this.liked
    },
    onComment() {
      this.comment = !this.comment
    },
  },
  created() {
    this.getPost(this.id).then(res => {
      this.data = res

      this.loaded = true

      if (this.data.atchType === 'photo') {
        this.getPhoto(this.data.atchId).then(res => {
          this.data.atchUrl = res.url
        })
      }
    })
  }
}
</script>

<style lang="scss" scoped>
@import "@/style.scss";

.post {
  margin-bottom: 12px;
  padding-bottom: $p4;
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
  margin-bottom: 8px;
}

.row {
  display: flex;
  // justify-content: flex-start;
  justify-content: space-between;
  align-items: center;
  margin: 0 $p4 0 $p4;
}
.btn-contaner {
  width: 50%;
  display: flex;
}

.btn {
  color: var(--fg);
  font-weight: 300;
  width: 50%;
  // width: max(25%, 100px);
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
