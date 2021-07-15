<template>
  <div class="header flex j-c a-c pointer" @click="click">
    <img src="../../assets/logo-color-dark.png" alt="Filecoin-logo" />
  </div>
</template>
<script>
export default {
  name: "Client",
  methods: {
    click() {
      this.go("Home");
    },
  },
};
</script>
<style lang='scss' scoped>
.header {
  padding: 10px 0;
  img {
    height: 80px;
  }
}
</style>
