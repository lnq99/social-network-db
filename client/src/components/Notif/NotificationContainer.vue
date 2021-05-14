<template>
  <scroll-container :items="notif">
    <template v-slot="slotProps">
      <link-card class="notif card-hl" :link="slotProps.item.link">
        {{ slotProps.item.text }}
      </link-card>
    </template>
  </scroll-container>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters({
      notif: 'notif/notif',
    }),
  },
}
</script>

<style scoped>
.notif {
  margin: 10px 20% 10px 20%;
}
</style>
