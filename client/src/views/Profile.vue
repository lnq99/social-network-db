<template>
  <div class="profile">
    <card class="col1">
      <el-affix :offset="-200">
        <card class="col-card card-hl">
          <avaname :name="name" :avatar="avatar"></avaname>
        </card>
        <card class="col-card card-hl"><photos></photos></card>
        <card class="col-card card-hl"><friends></friends></card>
      </el-affix>
    </card>
    <card class="col2">
      <card class="col-card card-hl">
        <intro></intro>
      </card>
      <post-container></post-container>
    </card>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Avaname from '@/components/Profile/AvaName.vue'
import Photos from '@/components/Profile/Photos.vue'
import Friends from '@/components/Profile/Friends.vue'
import Intro from '@/components/Profile/Intro.vue'
import PostContainer from '@/components/Post/PostContainer.vue'

export default {
  name: 'profile',
  components: { PostContainer, Avaname, Photos, Intro, Friends },
  computed: {
    ...mapState({ name: 'name', avatar: 'avatarL' }),
  },
  mounted() {
    console.log(this.name, this.avatar)
  }
}
</script>

<style scoped>
.profile {
  width: 100%;
  display: flex;
}
.col1 {
  flex: 1;
}
.col2 {
  flex: 2;
  padding-left: 12px;
}
.col-card {
  margin-bottom: 12px;
  padding: 18px;
}
</style>
