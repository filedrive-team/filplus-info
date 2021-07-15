<template>
  <div class="home bottom-20">
    <div class="chart-wrap">
      <div
          class="cell-warp"
          ref="chartLine"
          :style="{ width: '33%', height: '350px' }"
      ></div>
      <div
          class="cell-warp"
          ref="chartPie"
          :style="{ width: '33%', height: '350px' }"
      ></div>
      <div
          class="cell-warp"
          ref="chartPieRegion"
          :style="{ width: '33%', height: '350px' }"
      ></div>
    </div>
    <el-row justify="end" type="flex" class="top-5 bottom-20"></el-row>
    <div id="search">
      <search id="search_input"/>
    </div>
    <div id="table-border">
      <advanced-table
          :columns="columns"
          :transformData="transformData"
          ref="table"
          :api="getNotary"
      />
    </div>
  </div>
</template>

<script>
import {
  getNotary,
  getGrantedDaily,
  getProportionOfAllowance,
  getProportionOfAllowanceByLocation,
} from "@/api/notary.js";

import * as echarts from "echarts";
import dayjs from "dayjs";

export default {
  name: "Home",
  data(vm) {
    const formatter = (v) => {
      return v ? vm.unitConversion(v, 2) : "0 bytes";
    };
    return {
      columns: [
        {
          label: "Notary Name",
          key: "op",
          isComponent: true,
        },
        {
          label: "Organization",
          key: "organization",
        },
        {
          label: "Location",
          key: "location",
        },
        {
          label: "Total DataCap",
          key: "allowance",
          formatter,
        },
        {
          label: "Allocated",
          key: "grant_allowance",
          formatter,
        },
        {
          label: "Unallocated",
          key: "unallocated",
          formatter,
        },
      ],
      key: "",
      getNotary: getNotary,
      chartLine: null,
      getGrantedDaily: getGrantedDaily,
      getProportionOfAllowance: getProportionOfAllowance,
      getProportionOfAllowanceByLocation: getProportionOfAllowanceByLocation,
    };
  },
  mounted() {
    this.$nextTick(() => {
      this.drawLineChart();
      this.drawPieChart();
      this.drawPieRegionChart();
    });
  },
  methods: {
    transformData(list) {
      const vm = this;
      return list.map((item) => {
        const {notary_name, address, allowance, grant_allowance} = item;
        return {
          ...item,
          unallocated: Number(allowance) - Number(grant_allowance),
          op: {
            render() {
              return (
                  <div
                      onClick={() =>
                          vm.go("NotaryIndex", {addr: address, name: notary_name})
                      }
                      class="pointer"
                  >
                    <a style="color: DodgerBlue">{notary_name}</a>
                    <el-popover
                        title={notary_name}
                        trigger="hover"
                        content={address}
                    >
                    <span
                        slot="reference"
                        class="el-icon-warning-outline left-10"
                    />
                    </el-popover>
                  </div>
              );
            },
          },
        };
      });
    },
    async drawLineChart() {
      this.chartLine = echarts.init(this.$refs.chartLine);

      let option = {
        title: {
          text: "Datacap Allocation",
          textStyle: {
            fontSize: 14,
            align: "center",
          },
        },
        tooltip: {
          trigger: "axis",
        },
        calculable: true,
        xAxis: [
          {
            type: "category",
            boundaryGap: false,
            axisTick: {
              show: false,
            },
          },
        ],
        yAxis: [
          {
            type: "value",
            axisTick: {
              show: false,
            },
            name: "Tib",
          },
        ],
        series: [
          {
            name: "Allocated",
            type: "line",
          },
        ],
      };
      const res = await getGrantedDaily({limit: 7});
      const {code, data} = res;
      if (Number(code) === 0 && data) {
        if (Array.isArray(data) && data.length > 0) {
          let arrArr = [];
          let dimensions = Object.keys(data[0]);
          data.forEach((value, index) => {
            arrArr[index] = [];
            dimensions.forEach((val) => {
              let v = value[val];
              switch (val) {
                case "date":
                  v = dayjs.unix(v).format("MM/DD");
                  break;
                case "granted":
                  v = v / 1024 / 1024 / 1024 / 1024; // Tib
                  break;
              }
              arrArr[index].push(v || null);
            });
          });
          option.dataset = {
            source: arrArr,
          };
        }
      }

      this.chartLine.setOption(option);
    },
    async drawPieChart() {
      this.chartPie = echarts.init(this.$refs.chartPie);

      let option = {
        title: {
          text: "",
          left: "center",
          top: "center",
          textStyle: {
            fontSize: 14,
            align: "center",
          },
        },
        // graphic: {
        //   type: "text",
        //   left: "center",
        //   top: "40%",
        //   style: {
        //     text: "DataCap",
        //     textAlign: "center",
        //     fill: "#333",
        //     fontSize: 20,
        //     fontWeight: 700,
        //   },
        // },
        tooltip: {
          trigger: "item",
          formatter: "{b} {d}%",
        },
        calculable: true,
        series: [
          {
            name: "DataCap",
            type: "pie",
            radius: ["20%", "40%"],
            label: {
              formatter: "{b} {d}%",
              overflow: "break",
            },
            labelLine: {
              show: true,
            },
          },
        ],
      };
      const vm = this;
      const res = await getProportionOfAllowance();
      const {code, data} = res;
      if (Number(code) === 0 && data) {
        if (Array.isArray(data) && data.length > 0) {
          option.series[0].data = data;
          let total = 0;
          data.forEach((item) => {
            total += Number(item.value);
          });
          option.title.text = vm.unitConversion(total, 2);
        }
      }

      this.chartPie.setOption(option);
    },
    async drawPieRegionChart() {
      const vm = this;
      this.chartPieRegion = echarts.init(this.$refs.chartPieRegion);

      let option = {
        tooltip: {
          trigger: "item",
          formatter: function (params, ticket, callback) {
            let f =
                params.data.name +
                " " +
                vm.unitConversion(params.data.value, 2) +
                " (" +
                params.percent +
                "%)";
            callback(ticket, f);
          },
        },
        calculable: true,
        series: [
          {
            name: "Region ratio",
            type: "pie",
            radius: "60%",
            label: {
              formatter: "{b} {d}%",
            },
          },
        ],
      };
      const res = await getProportionOfAllowanceByLocation();
      const {code, data} = res;
      if (Number(code) === 0 && data) {
        if (Array.isArray(data) && data.length > 0) {
          option.series[0].data = data;
        }
      }

      this.chartPieRegion.setOption(option);
    },
  },
};
</script>

<style lang='scss' scoped>
#search {
  width: 100%;
  display: flex;
  justify-content: flex-end;
  margin-bottom: 20px;

}

#search_input {
  width: 20rem;
}

#table-border {
  border-top: rgb(229, 234, 242) 2px solid;
  border-right: rgb(229, 234, 242) 2px solid;
  border-bottom: rgb(229, 234, 242) 1px solid;
  border-left: rgb(229, 234, 242) 2px solid;
}

.chart-wrap {
  width: 100%;
  overflow: hidden;
  display: flex;
  flex-wrap: wrap;

  .cell-wrap {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: flex-start;
  }
}
</style>