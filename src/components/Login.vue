<template>
  <div>
    <el-button @click="loginVisible = true" icon="el-icon-user" circle />
    <el-dialog title="登录" :visible.sync="loginVisible" width="40%" center v-loading="loading">
      <div class="login-info">
        <el-input placeholder="用户名" v-model="username" clearable />
        <el-input
          placeholder="密码"
          v-model="password"
          show-password
          minlength="6"
          @keyup.enter.native="postLogin"
        />
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="loginVisible = false">取 消</el-button>
        <el-button type="primary" @click="postLogin">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data() {
    return {
      loginVisible: false,
      username: "",
      password: "",
      loading: false
    };
  },
  props: {},
  methods: {
    postLogin: function() {
      this.loading = true;
      axios
        .post(`/api/users/signin`, {
          Username: this.username,
          Password: this.password
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.msg);
          } else {
            setTimeout(() => {
              window.location.reload();
            }, 1000);
            this.$message({
              message: `${r.data.msg} 登录成功！`,
              type: "success"
            });
          }
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login-card {
  width: 300px;
}
.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}
.login-button {
  margin-top: 2px;
}
</style>
