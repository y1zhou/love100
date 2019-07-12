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
  computed: {
    tableData() {
      return this.$store.state.tableData;
    }
  }
};
</script>

<style>
.el-table .finished-row {
  text-decoration: line-through;
  font-style: italic;
  color: #c0c4cc;
}
</style>
