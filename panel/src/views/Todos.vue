<template>
  <div>
    <template v-for="(todo,i) in getTodoList">
      <todo-item :todo="todo" :index="i" :key="i" v-if="tag=='all'||todo.status===status" />
    </template>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
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
    };
  },
  computed: {
    ...mapGetters(["getTodoList"]),
  },
  created() {
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
  methods: {},
};
</script>

<style lang="scss" scoped>
</style>