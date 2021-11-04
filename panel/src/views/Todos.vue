<template>
  <div>
    <template v-for="(todo,i) in todoList">
      <todo-item :todo="todo.value" :index="i" :key="i" v-if="tag=='all'||todo.is_completed===status"/>
    </template>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import TodoItem from "../components/TodoItem.vue";

export default {
  name: "Todos",
  components: {
    TodoItem,
  },
  data() {
    return {
      tag: "all",
      status: "",
      todoList: []
    };
  },
  computed: {
    ...mapGetters(["getTodoList"]),
  },
  created() {
    this.getTodoList2()
    console.log(this.$route.params);
    console.log(this.tag);
  },
  beforeRouteUpdate(to, from, next) {
    this.tag = to.params.tag;
    switch (this.tag) {
      case "all":
        this.status = "";
        break;
      case "active":
        this.status = false;
        break;
      case "completed":
        this.status = true;
        break;
      default:
        break;
    }
    next();
  },
  methods: {
    getTodoList2() {
      const axios = require('axios');

// 向给定ID的用户发起请求
      const _this = this;
      axios.get('http://localhost:8080/v1/todos/0')
          .then(function (response) {
            // 处理成功情况
            console.log(response.data.result);
            _this.todoList = response.data.result
          })
    }
  },
};
</script>

<style lang="scss" scoped>
</style>