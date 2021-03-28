<template>
  <card class="p18 center hover">
    <router-link :to="link">
      <slot></slot>
    </router-link>
  </card>
</template>

<script>
export default {
  props: ['link'],
}
</script>

<style scoped>
a {
  color: #00796b;
  text-decoration: none;
  font-size: 1.2em;
  letter-spacing: 1px;
  word-spacing: 1px;
}
</style>
