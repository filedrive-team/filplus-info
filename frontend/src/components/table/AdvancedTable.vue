<template>
  <div class="advanced-table">
    <base-table
      :columns="columns"
      :data-source="dataSource"
      :show-pagination="true"
      :total="total"
      @page-change="handlePageChange"
      @size-change="handleSizeChange"
      :handleRowClick="handleRowClick"
      v-loading="loading"
      :emptyText="emptyText"
      v-bind="$attrs"
      v-on="$listeners"
    />
  </div>
</template>
<script>
export default {
  name: "AdvancedTable",
  props: {
    columns: {
      type: Array,
      default() {
        return [];
      },
    },

    api: {
      type: Function,
      default() {},
    },
    transformData: {
      type: Function,
      default: null,
    },
    filters: {
      type: Object,
      default() {
        return {};
      },
    },
    handleRowClick: {
      type: Function,
      default: null,
    },
  },
  data() {
    return {
      total: 0,
      pageSize: 20,
      page: 1,
      dataSource: [],
      loading: true,
      emptyText: "",
    };
  },
  computed: {
    pagination() {
      return {
        page: this.page,
        pageSize: this.pageSize,
      };
    },
    options() {
      const { page, pageSize } = this;
      return {
        page,
        pageSize,
        ...this.filters,
      };
    },
  },
  watch: {
    pagination() {
      this.getList(this.options);
    },
  },
  mounted() {
    this.getList(this.options);
  },
  methods: {
    handlePageChange(page) {
      this.page = page;
    },
    handleSizeChange(size) {
      this.page = 1;
      this.pageSize = size;
    },
    async getList(options) {
      let vm = this;
      const pageSize = options.pageSize;
      delete options.pageSize;
      const res = await this.api({
        ...options,
        page_size: pageSize,
      }).catch(() => {
        vm.dataSource = [];
        vm.total = 0;
        vm.emptyText = "Request data failed";
        vm.loading = false;
      });
      const { code, data } = res;
      if (Number(code) === 0 && data) {
        const { list, total } = data;
        const { transformData } = this;
        if (Array.isArray(list)) {
          this.dataSource = transformData ? transformData(list) : list;
          this.total = total;
        } else {
          this.dataSource = [];
          this.total = 0;
          this.emptyText = "Empty";
        }
      }
      vm.loading = false;
    },
    fresh() {
      this.page = 1;
      this.$nextTick(() => {
        this.getList(this.options);
      });
    },
  },
};
</script>
