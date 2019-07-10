<template>
  <el-container id="app">
    <el-header>
      <navbar />
    </el-header>
    <el-main class="main-container">
      <modify-table v-if="loggedIn" class="edit-table" />
      <display-table v-else class="main-table" />
      <el-backtop />
    </el-main>
  </el-container>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import DisplayTable from "./components/DisplayTable";
import ModifyTable from "./components/ModifyTable";
import axios from "axios";

export default {
  name: "app",
  components: {
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
          console.log(r.data.msg);
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
.edit-table {
  width: 100%;
  max-width: 1000px;
}
</style>
