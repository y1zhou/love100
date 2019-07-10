<template>
  <el-row type="flex" justify="center" class="header-bar">
    <el-col :span="6">
      <img src="../assets/love.png" />
    </el-col>
    <el-col :span="6">
      <div>这里记录了我们要一起做的100件事。</div>
    </el-col>
    <el-col :span="6">
      <el-button v-if="username != ''" @click="postLogout" icon="el-icon-switch-button" circle />
      <login v-else />
    </el-col>
  </el-row>
</template>

<script>
import Login from "./Login";
import axios from "axios";
export default {
  name: "Navbar",
  components: {
    Login
  },
  props: {},
  data() {
    return {
      showLogin: false,
      username: ""
    };
  },
  mounted() {
    axios
      .get("/api/user/", { withCredentials: true })
      .then(r => {
        if (r.data.err != "") {
          console.log(r.data.msg);
        } else {
          this.username = r.data.msg;
        }
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    postLogout: function() {
      axios
        .post("/api/user/signout", {
          withCredentials: true
        })
        .then(r => {
          if (r.data.err != "") {
            this.$message.error(r.data.err);
          } else {
            this.loggedIn = false;
            window.location.reload();
          }
        })
        .catch(err => {
          this.$message.error("You are not logged in.");
          console.log(err);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header-bar {
  align-items: center;
  justify-content: center;
}
.login-card {
  z-index: 10;
  position: fixed;
  margin-top: 5px;
}
</style>
