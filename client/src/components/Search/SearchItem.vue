<template>
  <card class="search-item row box-shadow">
    <div class="row">
      <link-card :link="{ name: 'Profile', params: { id: user.id } }">
        <short-info :id="user.id">
          <template v-slot="slotProps">
            <template class="row hover p18">
              <span>
                <el-avatar
                  class="ava"
                  :size="40"
                  :src="slotProps.avatars"
                ></el-avatar>
              </span>
              <div class="name">{{ slotProps.uname }}</div>
            </template>
          </template>
        </short-info>
      </link-card>
    </div>

    <template class="row" style="flex: 1">
      <div class="desc">{{ user.mutual }} mutual friends</div>

      <div class="tag-btn" v-if="user.type == 'friend'">
        <el-tag size="small" type="success">Friend</el-tag>
        <div>
          <el-dropdown size="mini">
            <span class="el-dropdown-link">
              More<i class="el-icon-arrow-down el-icon--right"></i>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>Unfriend</el-dropdown-item>
                <el-dropdown-item divided>Block</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
      <div class="tag-btn" v-else-if="user.type == 'request'">
        <el-tag size="small" type="warning">Request</el-tag>
        <div>
          <el-button class="btn" size="mini" type="primary">Accept</el-button>
          <el-button class="btn" size="mini" type="info">Delete</el-button>
        </div>
      </div>
      <div class="tag-btn" v-else-if="user.type == 'follow'">
        <el-tag size="small" type="info">Followed</el-tag>
        <el-button class="btn" size="mini" type="info" plain>Unsend</el-button>
      </div>
      <div class="tag-btn" v-else-if="user.type == 'block'">
        <el-tag size="small" type="danger">Block</el-tag>
        <el-button class="btn" size="mini" type="danger" plain
          >Unblock</el-button
        >
      </div>
      <div class="tag-btn" v-else>
        <span></span>
        <el-button class="" size="mini" type="success" plain
          >Send friend request</el-button
        >
      </div>
    </template>
  </card>
</template>

<script>
export default {
  props: ['user'],
}
</script>

<style scoped>
.row {
  display: flex;
}
.ava {
  margin: -10px 20px -10px -10px;
}
.name {
  width: 140px;
  text-align: left;
  white-space: nowrap;
  overflow: hidden;
}
.desc {
  margin-left: 24px;
  font-size: 0.8em;
  opacity: 0.7;
}
.search-item {
  margin-bottom: 18px;
  padding-right: 18px;
  /* justify-content: space-between; */
}
.btn {
  width: 80px;
}
.el-tag {
  margin-left: 10px;
  width: 64px;
  text-align: center;
  opacity: 0.8;
}
.tag-btn {
  display: flex;
  flex: 1;
  justify-content: space-between;
}
</style>
