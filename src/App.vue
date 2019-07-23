<template>
  <el-container id="app">
    <el-header>
      <navbar />
    </el-header>
    <el-main class="main-container" v-loading="isLoading">
      <modify-table v-if="loggedIn" />
      <display-table v-else />
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
      isLoading: true
    };
  },
  created() {
    // Check login status
    axios
      .get("/api/user/", { withCredentials: true })
      .then(r => {
        if (r.data.err != "") {
          console.log(r.data.msg);
        } else {
          this.$store.commit("login");
        }
      })
      .catch(err => {
        console.log(err);
      });
    // Get table from database
    axios
      .get(`/api/contents/`)
      .then(r => {
        if (r.data.err != "") {
          this.$message.error(r.data.msg);
        } else {
          console.log(`Found ${r.data.msg} entries.`);

          r.data.data.forEach(x => {
            x.CreatedAt = new Date(x.CreatedAt).toLocaleString("default", {
              hour12: false
            });
            x.UpdatedAt = new Date(x.UpdatedAt).toLocaleString("default", {
              hour12: false
            });
          });
          this.$store.commit("loadTable", r.data.data);
          this.isLoading = false;
        }
      })
      .catch(err => {
        console.log(err);
      });
  },
  computed: {
    loggedIn: function() {
      return this.$store.state.loggedIn;
    }
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
</style>
