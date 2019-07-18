<template>
  <div class="edit-table">
    <div class="floating-buttons">
      <el-tooltip effect="dark" content="添加" placement="top">
        <el-button @click="dialogFormVisible.add = true" icon="el-icon-plus" circle />
      </el-tooltip>
      <el-tooltip effect="dark" content="删除所有选中项" placement="top">
        <el-button @click="promptDeleteItems(delItems)" type="danger" icon="el-icon-delete" circle />
      </el-tooltip>
    </div>
    <el-table
      :data="tableData.filter(row => !search || row.Title.includes(search))"
      :row-class-name="tableRowClass"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="index" width="50"></el-table-column>
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="Title" label="项目" class-name="row-title"></el-table-column>
      <el-table-column prop="Comment" label="备注"></el-table-column>
      <el-table-column v-if="largeScreen" prop="CreatedAt" label="创建时间"></el-table-column>
      <el-table-column v-if="largeScreen" prop="UpdatedAt" label="更新时间"></el-table-column>

      <el-table-column fixed="right" min-width="150px">
        <template slot="header">
          <el-input v-model="search" size="mini" placeholder="输入关键字搜索" />
        </template>
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.Status"
            @change="toggleStatus(scope.$index)"
            style="margin-right: 10px;"
          />
          <el-tooltip effect="dark" content="编辑" placement="top">
            <el-button
              size="mini"
              icon="el-icon-edit"
              circle
              @click="promptEditItem(scope.$index, scope.row)"
            />
          </el-tooltip>
          <el-tooltip effect="dark" content="删除本项" placement="top">
            <el-button
              @click="promptDeleteItems([scope.row.ID])"
              size="mini"
              type="danger"
              icon="el-icon-delete"
              circle
            />
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      title="添加项目"
      :visible.sync="dialogFormVisible.add"
      @opened="$refs.addTitle.focus()"
      @closed="checkKeepAdding"
      :width="dialogWidth"
      center
    >
      <el-form :model="form" @submit.native.prevent>
        <el-form-item label="项目名称" label-width="80px">
          <el-input
            v-model="form.add.Title"
            ref="addTitle"
            clearable
            @keyup.enter.native="$refs.addComment.focus()"
          />
        </el-form-item>
        <el-form-item label="项目备注" label-width="80px">
          <el-input
            v-model="form.add.Comment"
            ref="addComment"
            clearable
            @keyup.enter.native="addItem"
          />
        </el-form-item>
        <el-form-item label="是否完成" label-width="80px">
          <el-switch v-model="form.add.Status" active-text="完成" inactive-text="未完成" />
        </el-form-item>
        <el-form-item label="继续添加" label-width="80px">
          <el-switch v-model="keepAdding" />
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="dialogFormVisible.add = false">取 消</el-button>
        <el-button type="primary" @click="addItem">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog
      title="编辑项目"
      :visible.sync="dialogFormVisible.edit"
      @opened="$refs.editTitle.select()"
      :width="dialogWidth"
      center
    >
      <el-form :model="form.edit" @submit.native.prevent>
        <el-form-item label="#" label-width="80px">{{ form.edit.ItemNum + 1 }}</el-form-item>
        <el-form-item label="项目名称" label-width="80px">
          <el-input
            v-model="form.edit.Title"
            ref="editTitle"
            clearable
            @keyup.enter.native="$refs.editComment.focus()"
          />
        </el-form-item>
        <el-form-item label="项目备注" label-width="80px">
          <el-input
            v-model="form.edit.Comment"
            ref="editComment"
            clearable
            @keyup.enter.native="updateItem(form.edit.ItemNum)"
          />
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="dialogFormVisible.edit = false">取 消</el-button>
        <el-button type="primary" @click="updateItem(form.edit.ItemNum)">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ModifyTable",
  data() {
    return {
      search: "",
      loading: false,
      keepAdding: false,
      delItems: [],
      dialogFormVisible: {
        add: false,
        edit: false
      },
      form: {
        add: {
          Title: "",
          Comment: "",
          Status: false
        },
        edit: {
          ItemNum: Number,
          ItemID: Number,
          Title: String,
          Comment: String,
          OldTitle: String,
          OldComment: String
        }
      }
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
    },
    handleSelectionChange(rows) {
      this.delItems = rows.map(x => x.ID);
    },
    addItem() {
      if (this.form.add.Title === "") {
        this.$message({
          message: "内容不能为空",
          type: "warning"
        });
        return;
      }
      axios
        .post("/api/content/", {
          Title: this.form.add.Title,
          Comment: this.form.add.Comment,
          ItemStatus: this.form.add.Status
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.msg);
          } else {
            this.$message({
              message: "添加内容成功！",
              type: "success"
            });
            r.data.data.CreatedAt = new Date(
              r.data.data.CreatedAt
            ).toLocaleString("default", {
              hour12: false
            });
            r.data.data.UpdatedAt = new Date(
              r.data.data.UpdatedAt
            ).toLocaleString("default", {
              hour12: false
            });

            this.$store.commit("addToTable", r.data.data);
            this.dialogFormVisible.add = false;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    checkKeepAdding() {
      if (this.form.add.Title != "" && this.keepAdding) {
        this.dialogFormVisible.add = true;
      }
      this.form.add.Title = "";
      this.form.add.Comment = "";
    },
    toggleStatus(index) {
      // The status changes with the @change event of el-switch
      axios
        .put("/api/content/status", {
          ItemID: this.tableData[index].ID,
          ItemStatus: this.tableData[index].Status
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.msg);
          } else {
            r.data.data.CreatedAt = new Date(
              r.data.data.CreatedAt
            ).toLocaleString("default", {
              hour12: false
            });
            r.data.data.UpdatedAt = new Date(
              r.data.data.UpdatedAt
            ).toLocaleString("default", {
              hour12: false
            });
            this.$store.commit("modifyTable", {
              index: index,
              row: r.data.data
            });
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    promptEditItem(index, row) {
      this.form.edit.ItemNum = index;
      this.form.edit.ItemID = row.ID;
      this.form.edit.OldTitle = row.Title;
      this.form.edit.Title = row.Title;
      this.form.edit.OldComment = row.Comment;
      this.form.edit.Comment = row.Comment;
      this.dialogFormVisible.edit = true;
    },
    updateItem(tableDataIndex) {
      if (this.form.edit.Title === "") {
        this.$message({
          message: "内容不能为空",
          type: "warning"
        });
        return;
      }
      if (
        this.form.edit.Title === this.form.edit.OldTitle &&
        this.form.edit.Comment === this.form.edit.OldComment
      ) {
        this.$message({
          message: "内容没有变化",
          type: "warning"
        });
        return;
      }
      axios
        .put("/api/content/", {
          ItemID: this.form.edit.ItemID,
          Title: this.form.edit.Title,
          Comment: this.form.edit.Comment
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.msg);
          } else {
            this.$message({
              message: "更新内容成功！",
              type: "success"
            });
            r.data.data.CreatedAt = new Date(
              r.data.data.CreatedAt
            ).toLocaleString("default", {
              hour12: false
            });
            r.data.data.UpdatedAt = new Date(
              r.data.data.UpdatedAt
            ).toLocaleString("default", {
              hour12: false
            });

            this.$store.commit("modifyTable", {
              index: tableDataIndex,
              row: r.data.data
            });
            this.dialogFormVisible.edit = false;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    promptDeleteItems(itemIDs) {
      this.$confirm(
        `此操作将删除 ${itemIDs.length} 条内容, 是否继续?`,
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      )
        .then(() => {
          this.deleteItems(itemIDs);
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    deleteItems(itemIDs) {
      if (itemIDs.length === 0) {
        this.$message({
          message: "没有选中需要删除的项目",
          type: "warning"
        });
        return;
      }
      axios
        .delete("/api/content/", {
          data: {
            ItemIDs: itemIDs
          }
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.msg);
          } else {
            this.$message({
              message: `成功删除 ${r.data.msg.length} 个项目`,
              type: "success"
            });
            this.$store.commit("deleteFromTable", { itemIDs: itemIDs });
          }
        })
        .catch(err => console.log(err));
    }
  },
  computed: {
    tableData() {
      return this.$store.state.tableData;
    },
    largeScreen() {
      return window.innerWidth > 900;
    },
    dialogWidth() {
      return this.largeScreen ? "40%" : "90%";
    }
  }
};
</script>

<style scoped>
.edit-table {
  align-items: center;
  width: 100%;
  max-width: 1300px;
}
.el-table .finished-row {
  color: #c0c4cc;
}
.el-table .finished-row .row-title {
  font-style: italic;
  text-decoration: line-through;
}
.floating-buttons {
  z-index: 10;
  position: fixed;
  left: 1rem;
  bottom: 1rem;
}
</style>
