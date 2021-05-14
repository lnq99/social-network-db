<template>
  <el-avatar class="avatar" :size="200" :src="avatar"></el-avatar>
  <h2 class="name">{{ name }}</h2>

  <link-card link="/"> News </link-card>
  <link-card link="/profile"> Profile </link-card>
  <link-card link="/photo"> Photo </link-card>
  <link-card link="/notif">
    <el-badge :value="100" :max="10" class="item"
      ><p>Notification</p>
    </el-badge>
  </link-card>
  <!-- <link-card link="/msg"> Messenger </link-card> -->
  <link-card link="/login"> Logout </link-card>
  <div class="center">
    <el-switch
      v-model="theme"
      class="switch"
      active-color="#111"
      inactive-color="#aaa"
      active-text="Dark"
      inactive-text="Light"
    >
    </el-switch>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    theme: {
      get() {
        return this.$store.state.isDark
      },
      set(value) {
        this.$store.dispatch('switchTheme', value)
      },
    },
    ...mapState({ name: 'name', avatar: 'avatarL' }),
  },
}
</script>

<style scoped>
.avatar {
  margin: 20px;
}
.name {
  margin-bottom: 4vh;
}
.switch {
  margin-top: 4vh;
  margin-left: -10px;
}
p {
  margin: 0 10px 0 10px;
}
</style>
