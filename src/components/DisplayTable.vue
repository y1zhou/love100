<template>
  <el-table :data="tableData" :row-class-name="tableRowClass">
    <el-table-column type="index" width="50"></el-table-column>
    <el-table-column prop="Title"></el-table-column>
  </el-table>
</template>

<script>
import axios from "axios";

export default {
  name: "DisplayTable",
  data() {
    return {
      tableData: []
    };
  },
  props: {},
  methods: {
    tableRowClass({ row }) {
      if (row.Status) {
        return "finished-row";
      } else {
        return "unfinished-row";
      }
    }
  },
  mounted() {
    axios
      .get(`/api/contents/`)
      .then(r => {
        if (r.data.err != "") {
          this.$message.error(r.data.msg);
        } else {
          console.log(`Found ${r.data.msg} entries.`);
          this.tableData = r.data.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }
};
</script>

<style>
.el-table .finished-row {
  text-decoration: line-through;
  color: #c0c4cc;
}
</style>
