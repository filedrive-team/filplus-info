<template>
  <div class="notary bottom-20">
    <el-row type="flex" class="top-20 bottom-20">
      <el-col :span="16" class="flex a-c">
        <div>
          <span
              @click="$router.go(-1)"
              class="pointer"
              style="color: DodgerBlue"
          >AllocationRecords /
          </span>
          <span>{{ client.name || "-" }} / </span>
          <span>{{ client.time }}</span>
        </div>
      </el-col>
      <!-- <el-col :span="8">
        <search @search='handleSearch' />
      </el-col> -->
    </el-row>

    <div id="table-border">
      <advanced-table
          :columns="columns"
          :transformData="transformData"
          ref="table"
          :api="getDeals"
          :filters="options"
      />
    </div>
  </div>
</template>

<script>
import {getDeals} from "@/api/notary.js";

export default {
  name: "Allocation",
  data(vm) {
    return {
      columns: [
        {
          label: "Epoch",
          key: "epoch",
        },
        {
          label: "Deal Id",
          key: "dealid_op",
          isComponent: true,
        },
        {
          label: "Payload Cid",
          key: "cid",
          formatter(v) {
            return vm.ellipsisByLength(v, 12);
          },
        },
        {
          label: "Provider",
          key: "provider",
        },
        {
          label: "Size",
          key: "piece_size",
          formatter(v) {
            return vm.unitConversion(v, 2);
          },
        },
        {
          label: "Start Epoch",
          key: "start_epoch",
        },
        {
          label: "End Epoch",
          key: "end_epoch",
        },
      ],
      getDeals,
      key: "",
    };
  },
  methods: {
    handleSearch(e) {
      console.log(e);
    },
    async api() {
      return {
        code: 0,
        message: "success",
        data: {
          list: [
            {
              epoch: 5656,
              deal: 1211,
              cid:
                  "babsbxsbjxb89suwhsjbnxsbxaxxmbxhiswjswi9swjskshhsbsbjskjsjskjjks",
              start: 100,
              end: 400,
              provider: "xnd",
              size: 123,
            },
          ],
          total: 98,
        },
      };
    },
    transformData(list) {
      return list.map((item) => {
        return {
          ...item,
          dealid_op: {
            render() {
              return (
                  <a
                      class="flex a-c j-c pointer"
                      style="color: DodgerBlue"
                      href={
                        "https://filscan.io/tipset/dsn-detail?dealid=" + item.dealid
                      }
                  >
                    {item.dealid}
                  </a>
              );
            },
          },
          op: {
            render() {
              return (
                  <div class="flex a-c j-c">
                    {item.time}
                    <span class="el-icon-caret-right left-10 font-20"/>
                  </div>
              );
            },
          },
        };
      });
    },
  },
  computed: {
    options() {
      const {client_address, start_epoch, allowance} = this.$route.query;
      return {
        client_address,
        start_epoch: Number(start_epoch),
        allowance,
      };
    },
    client() {
      const {name, time} = this.$route.query;
      return {
        name: name,
        time,
      };
    },
  },
};
</script>

<style>
#table-border {
  border-top: rgb(229, 234, 242) 2px solid;
  border-right: rgb(229, 234, 242) 2px solid;
  border-bottom: rgb(229, 234, 242) 1px solid;
  border-left: rgb(229, 234, 242) 2px solid;
}
</style>
