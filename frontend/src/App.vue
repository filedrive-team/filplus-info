<template>
  <div id="app">
    <Header />
    <div class="welcome flex j-c font-28 font-600 a-c">
      Welcome to the Filecoin Plus Dashboard
    </div>
    <div class="router-main">
      <router-view />
    </div>
    <Footer />
  </div>
</template>
<script>
import Header from "./components/header";
import Footer from "./components/footer";
export default {
  name: "App",
  components: {
    Header,
    Footer,
  },
};
</script>
<style lang="scss" scoped>
#app {
  display: flex;
  display: -webkit-flex; /* Safari */
  min-height: 100vh;
  flex-direction: column;

  .welcome {
    color: white;
    height: 100px;
    background: #3d8ff7;
  }
  .router-main {
    padding: 0 50px;
    flex: 1;
  }
}
</style>
