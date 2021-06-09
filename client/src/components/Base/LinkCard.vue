<template>
  <card :style="cls" class="center hover">
    <router-link :to="link">
      <slot></slot>
    </router-link>
  </card>
</template>

<script>
export default {
  props: ['link', 'cls'],
}
</script>