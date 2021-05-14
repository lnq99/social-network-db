<template>
  <div class="newsfeed">
    <new-post></new-post>
    <post-container></post-container>
  </div>
</template>

<script>
import Post from '@/components/Post/Post.vue'
import NewPost from '@/components/Post/NewPost.vue'
import PostContainer from '@/components/Post/PostContainer.vue'

export default {
  components: {
    Post,
    NewPost,
    PostContainer,
  },
}
</script>

<style scoped>
.newsfeed {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: var(--bg0);
  margin-left: 6vw;
  margin-right: 6vw;
}
</style>
