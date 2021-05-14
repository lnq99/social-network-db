<template>
  <scroll-container :items="posts">
    <template v-slot="slotProps">
      <post
        :key="slotProps.item.id"
        :id="slotProps.item.id"
        :name="slotProps.item.name"
        :text="slotProps.item.text"
        :photo="slotProps.item.photo"
        :time="slotProps.item.time"
        :avatar="slotProps.item.avatar"
      ></post>
    </template>
  </scroll-container>
</template>

<script>
import { mapGetters } from 'vuex'
import Post from './Post.vue'

export default {
  components: {
    Post,
  },
  computed: {
    ...mapGetters({ posts: 'newsfeed/posts' }),
  },
}
</script>
