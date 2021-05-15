<template>
  <div v-infinite-scroll="load" infinite-scroll-disabled="disabled">
    <div v-for="(item, index) in items">
      <slot :item="item"></slot>
    </div>
    <p v-if="loading">Loading...</p>
    <p v-if="noMore">No more</p>
  </div>
</template>

<script>
export default {
  props: ['items'],
  data() {
    return {
      count: this.items.length,
      loading: false,
    }
  },
  computed: {
    noMore() {
      return this.count >= this.items.length
    },
    disabled() {
      return this.loading || this.noMore
    },
  },
  methods: {
    load() {
      this.loading = true
      setTimeout(() => {
        this.count += 2
        this.loading = false
      }, 1000)
    },
  },
}
</script>
