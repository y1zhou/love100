<template>
  <div style="width: 100%">
    <div>
      <el-button :loading="loading" icon="el-icon-download" circle class="save-button" />
      <el-button type="danger" icon="el-icon-delete" circle class="delete-button" />
    </div>
    <el-table
      :data="tableData.filter(row => !search || row.Title.includes(search))"
      :row-class-name="tableRowClass"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="index" width="50"></el-table-column>
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="Title" label="Title"></el-table-column>
      <el-table-column prop="CreatedAt" label="Created At"></el-table-column>

      <el-table-column fixed="right">
        <template slot="header">
          <el-input v-model="search" size="mini" placeholder="输入关键字搜索" />
        </template>
        <template slot-scope="scope">
          <el-switch v-model="scope.row.Status" style="margin-right: 10px;" />
          <el-button
            size="mini"
            icon="el-icon-edit"
            circle
            @click="handleEdit(scope.$index, scope.row)"
          />
          <el-button size="mini" type="danger" icon="el-icon-delete" circle />
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ModifyTable",
  data() {
    return {
      loading: false,
      tableData: [],
      search: "",
      delItems: []
    };
  },
  props: {},
  computed: {
    itemsSelected: function() {
      return this.delItems.length > 0;
    }
  },
  methods: {
    tableRowClass({ row }) {
      if (row.Status) {
        return "finished-row";
      } else {
        return "unfinished-row";
      }
    },
    handleSelectionChange(rows) {
      this.delItems = rows.map(x => x.ID);
    },
    postItem() {}
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
          this.tableData.forEach(x => {
            x.CreatedAt = new Date(x.CreatedAt).toLocaleString("default", {
              hour12: false
            });
          });
        }
      })
      .catch(err => {
        console.log(err);
      });
  }
};
</script>

<style scoped>
.el-table .finished-row {
  text-decoration: line-through;
  font-style: italic;
  color: #c0c4cc;
}
.save-button {
  position: fixed;
  right: 100px;
  bottom: 150px;
}
.delete-button {
  position: fixed;
  right: 100px;
  bottom: 100px;
}
</style>
