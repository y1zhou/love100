<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="20">
        <img src="../assets/logo.png" />
      </el-col>
      <el-col :span="4">
        <el-button v-if="username != ''" @click="postLogout" icon="el-icon-user-solid" circle />
        <el-button v-else @click="showLogin = !showLogin" icon="el-icon-user" circle />
        <login v-if="showLogin" class="login-card" />
      </el-col>
    </el-row>
  </div>
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
          this.$message.error(r.data.err);
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
.login-card {
  z-index: 10;
  position: absolute;
  right: 5px;
  margin-top: 5px;
}
</style>
