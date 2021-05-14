<template>
  <div id="app" :class="this.$store.state.isDark ? 'dark' : 'light'">
    <router-view />
  </div>
</template>

<style lang="scss">
@import './style.scss';

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--text);
  background: var(--bg);
  height: 100%;
  width: 100%;
  display: flex;
}

.el-input__inner {
  border-radius: 100px !important;
}
.input .el-textarea__inner {
  border: none;
  background-color: var(--bg);
  color: var(--text);
  opacity: 0.8;
}
</style>
