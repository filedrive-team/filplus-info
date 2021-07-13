<script>
export default {
  name: 'Expand',
  props: {
    scope: {
      type: Object,
      default() {
        return {}
      }
    },
    component: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  render(h) {
    const component = this.component
    const scope = this.scope
    return h(component, {
      scope
    })
  }
}
</script>
