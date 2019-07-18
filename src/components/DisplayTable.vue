<template>
  <div class="main-table">
    <el-table :data="tableData" :row-class-name="tableRowClass">
      <el-table-column type="index" width="50"></el-table-column>
      <el-table-column prop="Title" label="要做的事" class-name="row-title"></el-table-column>
      <el-table-column prop="Comment" label="备注" class-name="row-comment"></el-table-column>
    </el-table>
  </div>
</template>

<script>
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
.main-table {
  width: 100%;
  max-width: 800px;
}
.el-table .finished-row {
  color: #c0c4cc;
}
.el-table .finished-row .row-title {
  font-style: italic;
  text-decoration: line-through;
}
.el-table .row-comment {
  color: #909399;
}
</style>
