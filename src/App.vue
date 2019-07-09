<template>
  <el-container id="app">
    <el-header height="210px">
      <navbar />
    </el-header>
    <el-main class="main-container">
      <display-table v-if="loggedIn" class="main-table" />
      <modify-table v-else class="main-table" />
    </el-main>
  </el-container>
</template>

<script>
import Login from "./components/Login.vue";
import Navbar from "./components/Navbar.vue";
import DisplayTable from "./components/DisplayTable";
import ModifyTable from "./components/ModifyTable";
import axios from "axios";

export default {
  name: "app",
  components: {
    Login,
    Navbar,
    DisplayTable,
    ModifyTable
  },
  data() {
    return {
      loggedIn: false
    };
  },
  mounted() {
    axios
      .get("/api/user/", { withCredentials: true })
      .then(r => {
        if (r.data.err != "") {
          this.$message.error(r.data.err);
        } else {
          this.loggedIn = true;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }
};
</script>

<style scoped>
#app {
  text-align: center;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}
.main-container {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 100%;
}
.main-table {
  width: 100%;
  max-width: 680px;
}
</style>
