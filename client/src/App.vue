<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #c9d1d9;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #c9d1d9;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

html,
body {
  height: 100%;
  padding: 0px;
  margin: 0px;
  color: #c9d1d9;
  background: #0d1117;
  scrollbar-color: #222 #0d1117;
}

::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.2);
}

::-webkit-scrollbar-thumb {
  border-radius: 8px;
  background-color: #222;
}

.el-card {
  /* color: #c9d1d9; */
  background: none !important;
  border: none !important;
  /* border: 1px solid #0d1117 !important; */
  /* border-radius: 10px !important; */
  box-shadow: none !important;
  /* box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important; */
}

.el-card__body {
  color: #c9d1d9;
  /* background: #161b22; */
  background: none !important;
  /* border: none !important; */
}
.el-card:hover {
  /* background: #161b22 !important; */
  /* filter: brightness(102%);
  box-shadow: #66666660 2px 2px 12px; */
}

hr {
  border-color: #222;
  margin: 10px 18px 10px 18px;
}

.row {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.hover:hover {
  background: #161b22 !important;
}

.p12 {
  padding: 12px;
}
.p18 {
  padding: 18px;
}
.center {
  text-align: center !important;
}
</style>
