<template>
  <div class="search">
    <h2>Search</h2>
    <el-row>
      <search-box @search="search"></search-box>
    </el-row>
    <hr />
    <br />
    <search-container :items="items"></search-container>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import SearchContainer from '../components/Search/SearchContainer.vue'
import SearchBox from '../components/Search/SearchBox.vue'

export default {
  components: { SearchContainer, SearchBox },
  data() {
    return {
      items: [],
      item: [
        {
          id: 1,
          type: 'friend',
          mutual: 8,
        },
        {
          id: 8,
          type: 'request',
          mutual: 2,
        },
        {
          id: 10,
          type: 'follow',
          mutual: 1,
        },
        {
          id: 12,
          type: 'block',
          mutual: 3,
        },
        {
          id: 11,
          type: '',
          mutual: 0,
        },
      ]
    }
  },
  methods: {
    ...mapActions({ searchProfile: 'profile/searchProfile' }),
    search(searchQuery) {
      if (searchQuery.length < 2) return
      this.searchProfile(searchQuery).then(res => {
        this.items = res
      })
    }
  }
}
</script>

<style scoped>
.search {
  margin-left: 8vw;
  margin-right: 8vw;
}
</style>