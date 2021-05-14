<template>
  <div class="profile">
    <card class="col1">
      <el-affix :offset="-200">
        <card class="col-card card-hl">
          <avaname></avaname>
        </card>
        <card class="col-card card-hl"><photos></photos></card>
        <card class="col-card card-hl"><friends></friends></card>
      </el-affix>
    </card>
    <card class="col2">
      <card class="col-card card-hl">
        <info></info>
      </card>
      <post-container></post-container>
    </card>
  </div>
</template>

<script>
import Avaname from '@/components/Profile/AvaName.vue'
import Photos from '@/components/Profile/Photos.vue'
import Friends from '@/components/Profile/Friends.vue'
import Info from '@/components/Profile/Info.vue'
import PostContainer from '@/components/Post/PostContainer.vue'

export default {
  name: 'profile',
  components: { PostContainer, Avaname, Photos, Info, Friends },
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
