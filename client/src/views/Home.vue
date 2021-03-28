<template>
  <div class="home">
    <el-container style="height: 100vh">
      <left-panel></left-panel>
      <main-view></main-view>
      <right-panel></right-panel>
    </el-container>
  </div>
</template>

<script>
// @ is an alias to /src
import MainView from '@/components/MainView.vue'
import LeftPanel from '@/components/LeftPanel.vue'
import RightPanel from '@/components/RightPanel.vue'

export default {
  name: 'Home',
  components: {
    MainView,
    LeftPanel,
    RightPanel,
  },
}
</script>
