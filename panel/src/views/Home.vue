<template>
  <div id="app">
    <div id="main">
      <div id="title">todos</div>
      <div id="section">
        <Edit />
        <Todo v-if="$store.state.todoList.length>0" />
      </div>
    </div>
  </div>
</template>

<script>
import Edit from "../components/Edit";
import Todo from "../components/Todo";
export default {
  name: "App",
  components: {
    Edit,
    Todo,
  },
  created() {
    if (window.localStorage.getItem("todos")) {
      this.$store.dispatch(
        "changeTodoList",
        JSON.parse(window.localStorage.getItem("todos"))
      );
    }
  },
};
</script>

<style lang="scss" scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  display: flex;
  justify-content: center;
  #main {
    width: 550px;
    #title {
      font-size: 100px;
      font-weight: 100;
      color: #ead7d7;
    }
    #section {
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
    }
  }
}
</style>
