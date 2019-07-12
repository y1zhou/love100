<template>
  <el-row type="flex" justify="center" class="header-bar">
    <el-col :span="3">
      <img src="../assets/love.png" />
    </el-col>
    <el-col :span="12" :offset="3">
      <div>这里记录了我们要一起做的100件事。</div>
      <el-progress
        :text-inside="true"
        :stroke-width="15"
        :percentage="percentFinished"
        color="#ff3366"
      />
    </el-col>
    <el-col :span="3" :offset="3">
      <el-tooltip v-if="showLogout" effect="dark" content="登出" placement="bottom">
        <el-button @click="postLogout" icon="el-icon-switch-button" circle />
      </el-tooltip>
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
            this.$store.commit("logout");
            window.location.reload();
          }
        })
        .catch(err => {
          this.$message.error("You are not logged in.");
          console.log(err);
        });
    }
  },
  computed: {
    showLogout() {
      return this.$store.state.loggedIn;
    },
    // Calculate percentage of finished entries
    percentFinished() {
      return this.$store.getters.finishedPercentage;
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
