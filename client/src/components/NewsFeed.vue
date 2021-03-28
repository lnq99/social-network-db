<template>
  <div
    class="list newsfeed"
    style="overflow:auto"
    v-infinite-scroll="load"
    infinite-scroll-disabled="disabled"
  >
    <div v-for="i in count">
      <post
        :key="posts[0].id"
        :id="posts[0].id"
        :name="posts[0].name"
        :text="posts[0].text"
        :photo="posts[0].photo"
        :time="posts[0].time"
        :avatar="posts[0].avatar"
      ></post>
    </div>
    <p v-if="loading">Loading...</p>
    <p v-if="noMore">No more</p>
  </div>
</template>

<script>
import posts from './post.js'
import Post from './Post.vue'

export default {
  components: {
    Post,
  },
  data() {
    return {
      posts,
      count: 10,
      loading: false,
    }
  },
  computed: {
    noMore() {
      return this.count >= 20
    },
    disabled() {
      return this.loading || this.noMore
    },
  },
  methods: {
    load() {
      this.loading = true
      setTimeout(() => {
        this.count += 2
        this.loading = false
      }, 1000)
    },
  },
}
</script>

<style scoped>
.newsfeed {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-left: 10vw;
  padding-right: 10vw;
}
</style>
