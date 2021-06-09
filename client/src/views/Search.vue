<template>
  <div class="search">
    <h2>Search</h2>
    <!-- <search></search> -->
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
          name: 'Quang',
          avtars: 'http://placehold.it/40x40',
          type: 'friend',
          mutualFriends: 8,
        },
        {
          id: 8,
          name: 'John',
          avtars: 'http://placehold.it/40x40',
          type: 'request',
          mutualFriends: 2,
        },
        {
          id: 10,
          name: 'Robert',
          avtars: 'http://placehold.it/40x40',
          type: 'follow',
          mutualFriends: 1,
        },
        {
          id: 10,
          name: 'Arina',
          avtars: 'http://placehold.it/40x40',
          type: 'block',
          mutualFriends: 3,
        },
        {
          id: 10,
          name: 'Eliza',
          avtars: 'http://placehold.it/40x40',
          type: '',
          mutualFriends: 0,
        },
      ]
    }
  },
  methods: {
    ...mapActions({ searchProfile: 'profile/searchProfile' }),
    search(searchQuery) {
      if (searchQuery.length < 2) return
      console.log(searchQuery)
      this.searchProfile(searchQuery).then(res => {
        this.items = res
        console.log(res)
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