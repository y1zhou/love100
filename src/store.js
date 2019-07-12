import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tableData: [],
    loggedIn: false
  },
  getters: {
    finishedPercentage: state => {
      return state.tableData.length > 0
        ? Math.round(
            (100 * state.tableData.filter(x => x.Status).length) /
              state.tableData.length
          )
        : 0
    }
  },
  mutations: {
    login(state) {
      state.loggedIn = true
    },
    logout(state) {
      state.loggedIn = false
    },
    loadTable(state, table) {
      state.tableData = table
    },
    addToTable(state, row) {
      state.tableData.push(row)
    },
    modifyTable(state, payload) {
      state.tableData.splice(payload.index, 1, payload.row)
    },
    deleteFromTable(state, payload) {
      state.tableData = state.tableData.filter(
        x => payload.itemIDs.indexOf(x.ID) === -1
      )
    }
  },
  actions: {}
})
