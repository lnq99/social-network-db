<template>
  <card class="center hover">
    <router-link :to="link">
      <div class="p18">
        <slot></slot>
      </div>
    </router-link>
  </card>
</template>

<script>
export default {
  props: ['link'],
}
</script>

<style scoped>
a > div {
  font-size: 1.2em;
  font-weight: 200;
  letter-spacing: 1px;
  word-spacing: 1px;
}
</style>
