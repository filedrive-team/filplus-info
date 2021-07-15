<template>
  <div class="base-table">
    <slot name="header"/>
    <el-table
        ref="table"
        :data="dataSource"
        v-bind="$attrs"
        v-on="$listeners"
        @sort-change="handleSortChange"
        @row-click="rowClick"
    >
      <el-table-column
          v-for="item in columns"
          :key="item.key"
          :label="item.label"
          :sortable="item.sortable ? 'custom' : false"
          :prop="item.key"
          :align="item.align || 'center'"
          :width="item.width"
          :summary-method="getSummaries"
          :show-summary="showSummary"
      >

        <template slot-scope="scope">
          <span v-if="!item.isComponent"
          >{{
              item.formatter
                  ? item.formatter(scope.row[item.key], scope.row)
                  : scope.row[item.key]
            }}{{ ` ${item.unit || ""}` }}</span
          >
          <component :is="scope.row[item.key]" v-else/>
        </template>
      </el-table-column>
      <el-table-column
          v-if="expand"
          type="expand"
          :width="hideExpand ? '1px' : '20px'"
      >
        <template slot-scope="scope">
          <expand :scope="scope" :component="component"/>
        </template>
      </el-table-column>
      <div slot="empty">
        <p>{{ emptyText }}</p>
      </div>
    </el-table>
    <el-pagination
        v-if="showPagination && dataSource.length"
        :current-page="currentPage"
        hide-on-single-page
        :page-sizes="pageSizes"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />
    <slot name="footer"/>
  </div>
</template>
<script>
import Expand from "./Expand";

export default {
  name: "BaseTable",
  components: {
    Expand,
  },
  inheritAttrs: false,
  props: {
    component: {
      type: Object,
      default() {
        return {
          functional: true,
          render() {
            return null;
          },
        };
      },
    },
    expand: {
      type: Boolean,
      default: false,
    },
    dataSource: {
      type: Array,
      default() {
        return [];
      },
    },
    columns: {
      type: Array,
      default() {
        return [];
      },
    },
    showPagination: {
      type: Boolean,
      default: false,
    },
    total: {
      type: Number,
      default: 0,
    },
    pageSizes: {
      type: Array,
      default() {
        return [20, 50, 100, 200];
      },
    },
    pageSize: {
      type: Number,
      default: 20,
    },
    currentPage: {
      type: Number,
      default: 1,
    },
    defaultSort: {
      type: Object,
      default() {
        return {};
      },
    },
    showSummary: {
      type: Boolean,
      default: false,
    },
    hideExpand: {
      type: Boolean,
      default: false,
    },
    handleRowClick: {
      type: Function,
      default: null,
    },
    emptyText: {
      type: String,
      default: "",
    },
  },
  mounted() {
    this.$emit("getHandle", this.$refs.table);
  },
  methods: {
    handleSizeChange(size) {
      this.$emit("size-change", size);
    },
    handleCurrentChange(num) {
      this.$emit("page-change", num);
    },
    handleSortChange(v) {
      this.$emit("sort-change", v);
    },
    getSummaries(v) {
      return this.$emit("get-summary", v);
    },
    rowClick(row) {
      const {handleRowClick} = this;
      handleRowClick ? handleRowClick(row) : null;
    },
  },
};
</script>
<style lang="scss" scoped>
.base-table {
  & ::v-deep .el-pagination {
    text-align: center;
    margin-top: 20px;
  }

  & ::v-deep .el-table__row.expanded {
    td {
      border-bottom-color: transparent;
    }
  }
}
</style>

