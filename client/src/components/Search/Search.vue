<template>
  <div class="ui cards dropdown" style="margin: 10px; width: 100%">
    <el-input
      placeholder="Search..."
      suffix-icon="el-icon-search"
      v-model="searchQuery"
    />
    <div style="position: absolute; z-index: 1;">
      <div
        class="content dropdown-content"
        v-for="user in searchedUsers"
        :key="user.name"
      >
        {{ user.name }}<br />
        {{ user.phone }}
      </div>
    </div>
  </div>
</template>

<script>
import { computed, reactive, ref } from 'vue'
export default {
  setup() {
    const users = reactive([
      {
        name: 'Delaney French',
        email: 'delaneyfrench@imkan.com',
        phone: '+1 (973) 566-3049',
      },
      {
        name: 'Salinas Rowland',
        email: 'salinasrowland@imkan.com',
        phone: '+1 (866) 586-3960',
      },
    ])
    const searchQuery = ref('')
    const searchedUsers = computed(() => {
      return users.filter((user) => {
        if (searchQuery.value === '') return false
        return (
          user.name.toLowerCase().indexOf(searchQuery.value.toLowerCase()) != -1
        )
      })
    })
    return { searchedUsers, searchQuery }
  },
}
</script>

<style lang="scss" scoped>
.content {
  border-top: 1px solid #eee;
  padding: 16px;
  background: var(--hl);
}
.dropdown {
  position: relative;
  display: inline-block;
}

.dropdown-content {
  display: none;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  padding: 12px 16px;
  z-index: 1;
}

.dropdown:hover .dropdown-content {
  display: block;
}
</style>
