<template>
  <el-input
    :placeholder="placeholder"
    prefix-icon="el-icon-search"
    @keydown.native.enter="handleSearch"
    v-model="key"
  />
</template>
<script>
export default {
  name: "Search",
  data() {
    return {
      key: "",
    };
  },
  props: {
    placeholder: {
      type: String,
      default: "Client Name / Filecoin Address",
    },
  },
  methods: {
    handleSearch() {
      this.$emit("search", this.key);
      if (this.key != "") {
        let ok = /^f[1-3].*/g.test(this.key);
        if (ok) {
          // filecoin address
        } else {
          // client name
          this.go("Client", {
            client_name: this.key,
          });
        }
      }
    },
  },
};
</script>

