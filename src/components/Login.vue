<template>
  <el-card class="login-card" shadow="always" v-loading="loading">
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
    <div class="login-button">
      <el-button @click="postLogin">登录</el-button>
    </div>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data() {
    return {
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
            this.$message({
              message: `${r.data.msg} 登录成功！`,
              type: "success"
            });
            setTimeout(() => {
              window.location.reload();
            }, 1500);
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
