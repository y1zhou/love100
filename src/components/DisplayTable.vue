<template>
  <div id="main-table">
    <el-table :data="tableData" :row-class-name="tableRowClass" class="display-table">
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
      let res = [];
      if (row.Status) {
        res.push("finished-row");
      } else {
        res.push("unfinished-row");
      }
      res.push(this.userGender(row.AuthorID));
      return res.join(" ");
    },
    userGender(userID) {
      return this.$store.state.userData
        .filter(x => x.ID == userID)
        .map(x => x.Gender) == "F"
        ? "girl-row"
        : "guy-row";
    }
  },
  computed: {
    tableData() {
      return this.$store.state.tableData;
    }
  }
};
</script>

<style lang="scss">
#main-table {
  width: 100%;
  max-width: 800px;
}
.display-table {
  color: #606266;
  .finished-row {
    color: #c0c4cc;
    .row-title {
      font-style: italic;
      text-decoration: line-through;
    }
  }
  .guy-row {
    .row-comment {
      color: #909399;
    }
  }
  .girl-row {
    .row-comment {
      color: #eb83ab;
    }
  }
}
</style>
