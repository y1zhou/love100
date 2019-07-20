<template>
  <el-row type="flex" justify="center" class="header-bar">
    <el-col :span="3">
      <img src="../assets/love.png" />
    </el-col>
    <el-col :span="12" :offset="3">
      <div>这里记录了我们要一起做的 {{ $store.getters.numEntries }} 件事。</div>
      <el-progress
        :text-inside="true"
        :stroke-width="15"
        :percentage="percentFinished"
        color="#ff3366"
      />
    </el-col>
    <el-col :span="3" :offset="3">
      <login />
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
  computed: {
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
</style>
