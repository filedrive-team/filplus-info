<template>
  <div class="notary">
    <el-row type="flex" class="top-20 bottom-20">
      <el-col :span="16" class="flex a-c">
        <div>
          <span
              @click="$router.go(-1)"
              class="pointer"
              style="color: DodgerBlue"
          >HomePage /
          </span>
          <span>{{ navigation.notary_name }} </span>
        </div>
      </el-col>
    </el-row>
    <!-- <el-row
      justify="end"
      type="flex"
      class="top-20 bottom-20"
    >
      <el-col :span="8">
        <search @search='handleSearch' />
      </el-col>
    </el-row> -->
    <div id="table-border">
      <advanced-table
          :columns="columns"
          :transformData="transformData"
          ref="table"
          :api="getAllocated"
          :filters="options"
      />
    </div>
  </div>
</template>

<script>
import {getAllocated} from "@/api/notary.js";
import dayjs from "dayjs";

export default {
  name: "Notary",
  data(vm) {
    return {
      key: "",
      columns: [
        {
          label: "Client Name",
          key: "client_name_op",
          isComponent: true,
        },
        {
          label: "Filecoin Address",
          key: "client_op",
          isComponent: true,
        },
        {
          label: "DataCap Allocated",
          key: "allowance",
          formatter(v) {
            return vm.unitConversion(v, 2);
          },
        },
        {
          label: "Time",
          key: "op",
          isComponent: true,
        },
      ],
      getAllocated,
    };
  },
  computed: {
    options() {
      const address = this.$route.query.addr;
      return {
        address,
        client_address: this.key,
      };
    },
    navigation() {
      const name = this.$route.query.name;
      return {
        notary_name: name,
      };
    },
  },
  methods: {
    handleSearch(e) {
      this.key = e;
      this.$refs.table.fresh();
    },
    async api() {
      return {
        code: 0,
        message: "success",
        data: {
          list: [
            {
              client: "FileDrive ",
              address:
                  "f3u5nfw77pc7l2dsnlbzxz3jcafyjymgbkxbp4vifaoojcvhzv3sgroez6z7rrlxh5z3xpuw4pcnwiks6xvyna",
              time: "2021/02/02 19:00:30",
              allocated: "90TiB",
            },
          ],
          total: 98,
        },
      };
    },
    transformData(list) {
      const vm = this;
      return list
          .filter((item) => item)
          .map((item) => {
            return {
              ...item,
              client_name_op: {
                render() {
                  return (
                      <div
                          class="flex a-c j-c pointer"
                          style="color: DodgerBlue"
                          onClick={() =>
                              vm.go("Client", {
                                client_name: item.client_name,
                              })
                          }
                      >
                        {item.client_name}
                      </div>
                  );
                },
              },
              client_op: {
                render() {
                  return (
                      <div class="flex a-c j-c">
                        {vm.ellipsisByLength(item.client, 12)}
                        <span
                            class="el-icon-document-copy left-10 pointer address"
                            data-clipboard-text={item.client}
                            onClick={() => {
                              var clipboard = new this.Clipboard(".address");
                              clipboard.on("success", () => {
                                // 释放内存
                                clipboard.destroy();
                              });
                              clipboard.on("error", () => {
                                // 不支持复制
                                console.log(
                                    "The browser does not support automatic replication."
                                );
                                // 释放内存
                                clipboard.destroy();
                              });
                            }}
                        ></span>
                      </div>
                  );
                },
              },
              op: {
                render() {
                  return (
                      <div class="flex a-c j-c">
                        {dayjs.unix(item.block_time).format("YYYY-MM-DD HH:mm:ss")}
                        <span
                            class="el-icon-caret-right left-10 font-20 pointer"
                            style="color: DodgerBlue"
                            onClick={() =>
                                vm.go("Allocation", {
                                  client_address: item.client,
                                  name: item.client_name,
                                  start_epoch: item.epoch,
                                  allowance: item.allowance,
                                  time: dayjs
                                      .unix(item.block_time)
                                      .format("YYYY-MM-DD HH:mm:ss"),
                                })
                            }
                        />
                      </div>
                  );
                },
              },
            };
          });
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
