<template>
  <div style="z-index=999">
    <div class="input-container">
      <a-icon
        type="down"
        class="icon"
        :style="getStatus?'color:black;':''"
        v-if="$store.state.todoList.length>0"
        @click="clickIcon"
      />
      <a-input
        class="input"
        placeholder="What needs to be done?"
        size="large"
        v-model="todo"
        @pressEnter="addTodoList"
      ></a-input>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";
export default {
  name: "Edit",
  data() {
    return {
      todo: "",
    };
  },
  computed: {
    ...mapGetters(["getTodoList", "getStatus"]),
  },
    created(){
      if(this.getTodoList.length>0){
          this.$router.push({path:'/all'})
      }
  },
  methods: {
    ...mapMutations(["changeStatus", "changeTodoList"]),
    addTodoList() {
      if (this.todo.trim() != "") {
        this.$store.dispatch("addTodoList", { name: this.todo, status: false });
        this.todo = "";
        if (this.$route.path == "/") {
          this.$router.push({ path: "/all" });
        }
      }
    },
    clickIcon() {
      this.changeStatus(!this.getStatus);
      for (let i = 0; i < this.getTodoList.length; i++) {
        this.getTodoList[i].status = this.getStatus;
      }
      this.changeTodoList(this.getTodoList);
    },
  },
};
</script>

<style lang='scss' scoped>
.input-container {
  padding: 10px;
  border: 1px solid #fdfdfd;
  background-color: #fefefe;
  display: flex;
  justify-content: space-around;
  align-content: center;
  .icon {
    position: relative;
    top: 10px;
    margin: 0 10px;
    color: #e8e7e6;
    font-size: 20px;
    cursor: pointer;
  }
}
.ant-input {
  border: none;
}
.ant-input:focus {
  box-shadow: none;
}
</style>