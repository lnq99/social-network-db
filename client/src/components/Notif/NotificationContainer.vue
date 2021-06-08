<template>
  <scroll-container :items="notif">
    <template v-slot="slotProps">
      <notif-item :notif="slotProps.item"></notif-item>
    </template>
  </scroll-container>
</template>

<script>
import { mapGetters } from 'vuex'
import LinkCard from '../Base/LinkCard.vue'
import NotifItem from './NotifItem.vue'

export default {
  components: { LinkCard, NotifItem },
  computed: {
    ...mapGetters({ notif: 'notif/notif' }),
  },
  created() {
    this.$store.dispatch('notif/getNotif')
    console.log(this.notif)
  }
}
</script>

<style scoped>
.notif {
  margin: 10px 20% 10px 20%;
}
</style>
