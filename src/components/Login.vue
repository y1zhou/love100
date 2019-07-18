<template>
  <div>
    <el-tooltip effect="dark" content="登录" placement="bottom">
      <el-button @click="loginVisible = true" icon="el-icon-user" circle />
    </el-tooltip>
    <el-dialog
      title="登录"
      :visible.sync="loginVisible"
      @opened="$refs.username.focus()"
      :width="loginCardWidth"
      center
      v-loading="loading"
    >
      <div>
        <el-input ref="username" placeholder="用户名" v-model="username" clearable />
        <el-input
          placeholder="密码"
          v-model="password"
          show-password
          minlength="6"
          @keyup.enter.native="postLogin"
        />
      </div>
      <span slot="footer">
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
      if (this.username.length === 0) {
        this.$message.error("用户名为空");
        this.$refs.username.focus();
      } else if (this.password.length < 6) {
        this.$message.error("密码长度过短");
      } else {
        this.loading = true;
        axios
          .post(`/api/users/signin`, {
            Username: this.username,
            Password: this.password
          })
          .then(r => {
            if (r.data.err != "") {
              this.loading = false;
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
            this.$message.error("登录错误，请重试");
            this.loading = false;
            this.loginVisible = false;
            console.log(err);
          });
      }
    }
  },
  computed: {
    loginCardWidth() {
      return window.innerWidth > 900 ? "30%" : "90%";
    }
  }
};
</script>
