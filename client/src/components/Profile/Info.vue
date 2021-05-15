<template>
  <h2>
    Intro
    <el-button class="btn-right" @click="isEdit = true">
      Edit
    </el-button>
  </h2>
  <p v-for="i in info.split('\n')">{{ i }}</p>
  <el-dialog title="Edit Info" v-model="isEdit" width="40%" center>
    <el-input
      type="textarea"
      :autosize="{ minRows: 2, maxRows: 4 }"
      v-model="infoContent"
    >
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="onSaveInfo">Save</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import { ref } from 'vue'
import { mapGetters, mapActions } from 'vuex'

export default {
  data() {
    return {
      isEdit: false,
      infoContent: ref(''),
    }
  },
  computed: {
    ...mapGetters(['info']),
  },
  methods: {
    ...mapActions(['saveInfo']),
    onSaveInfo() {
      this.saveInfo(this.infoContent)
      this.isEdit = false
    },
  },
  created() {
    if (this.info.length < 10) {
      this.infoContent = 'Works at ...\nStudied at \nLives in ...\nFrom ...'
    } else {
      this.infoContent = this.info
    }
  },
}
</script>

<style scoped>
.btn-right {
  float: right;
  margin-top: -6px;
  padding-top: 4px;
  padding-bottom: 4px;
}
</style>
