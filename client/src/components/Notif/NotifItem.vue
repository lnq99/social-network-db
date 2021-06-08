<template>
  <short-info :id="notif.fromUserId">
    <template v-slot="slotProps">
      <template v-if="notif.type == 'request'">
        <link-card class="notif card-hl" :link="`/profile/${notif.fromUserId}`">
          {{ slotProps.uname }} send you a friend request.
        </link-card>
      </template>
      <template v-else-if="notif.type == 'accept'">
        <link-card class="notif card-hl" :link="`/profile/${notif.fromUserId}`">
          {{ slotProps.uname }} accepted your friend request.
        </link-card>
      </template>
      <template v-else-if="notif.type == 'cmt'">
        <link-card class="notif card-hl" :link="`/post/${notif.postId}`">
          {{ slotProps.uname }} commented on your post.
        </link-card>
      </template>
      <template v-else-if="notif.type == 'like'">
        <link-card class="notif card-hl" :link="`/post/${notif.postId}`">
          {{ slotProps.uname }} liked your post.
        </link-card>
      </template>
    </template>
  </short-info>
</template>

<script>
import LinkCard from '../Base/LinkCard.vue'
import ShortInfo from '../Profile/ShortInfo.vue'


export default {
  components: { LinkCard, ShortInfo },
  props: ['notif'],
}
</script>

<style scoped>
.notif {
  margin: 10px 20% 10px 20%;
}
</style>
