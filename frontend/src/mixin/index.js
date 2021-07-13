const mixin = {
  methods: {
    ellipsisByLength(str = "", length = 6, flag = true) {
      return str.length > 2 * length && flag
        ? `${str.slice(0, length)}...${str.slice(-length)}`
        : str
    },
    go(name, query = {}) {
      this.$router.push({
        name,
        query
      })
    },
    unitConversion(item, len) {
      let positive = true
      if (item == 0) {
        return "0 bytes"
      }
      if (item < 0) {
        positive = false
        item = Math.abs(item)
      }
      let k = 1024
      let sizes = [
        "bytes",
        "KiB",
        "MiB",
        "GiB",
        "TiB",
        "PiB",
        "EiB",
        "ZiB",
        "YiB"
      ]
      let c = Math.floor(Math.log(item) / Math.log(k))
      if (c < 0) {
        item = 0
      } else {
        item = (item / Math.pow(k, c)).toFixed(len) + " " + sizes[c]
      }

      return positive ? item : `-${item}`
    },


  }
}
export default mixin