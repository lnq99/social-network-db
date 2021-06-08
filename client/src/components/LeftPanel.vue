<template>
  <el-avatar class="ava" :size="200" :src="avatar"></el-avatar>
  <h2 class="name">{{ name }}</h2>

  <link-card link="/"> News </link-card>
  <link-card :link="{ name: 'Profile', params: { id } }"> Profile </link-card>
  <link-card :link="{ name: 'Photo', params: { id } }"> Photo </link-card>
  <link-card link="/notif">
    <el-badge :value="100" :max="10" class="item"
      ><p>Notification</p>
    </el-badge>
  </link-card>
  <!-- <link-card link="/msg"> Messenger </link-card> -->
  <link-card link="/logout"> Logout </link-card>
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
    ...mapState({ id: 'id', name: 'name', avatar: 'avatarl' }),
  },
}
</script>

<style scoped>
.ava {
  margin-top: 20px;
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
