<template>
  <div class="todo">
    <a-checkbox class="checkbox" :checked="todo.status" @change="checkboxChange" />
    <div class="todo-name" @dblclick="edit">
      <a-input
        v-if="editMode==true"
        size="large"
        v-model="todo"
        @pressEnter="changeTodo"
        v-autoFocus
        @blur="blurEdit"
      />
      <span v-else :style="todo.status?'text-decoration:line-through;color:#d9d9d9':''">{{todo}}</span>
    </div>
    <a-icon type="close" class="icon" @click="delTodo(index)" />
  </div>
</template>


<script>
import { mapGetters, mapActions } from "vuex";
export default {
  name: "TodoItem",
  props: {
    todo: {
      type: String,
    },
    index: {
      type: Number,
    },
  },
  data() {
    return {
      editMode: false,
      value: this.todo.name,
    };
  },
  computed: {
    ...mapGetters(["getTodoList", "getLeft"]),
  },
  mounted(){
      this.$nextTick(()=>{
          document.addEventListener('keyup',(e)=>{
              if(e.keyCode==27){
                  this.todo.name = this.value;
                  this.editMode=false;
              }
          })
      })
  },
  directives: {
    autoFocus: {
      inserted: function (el) {
        el.focus();
      },
    },
  },
  methods: {
    ...mapActions(["delTodo", "changeStatus", "changeTodoList"]),
    checkboxChange() {
      this.getTodoList[this.index].status = !this.getTodoList[this.index]
        .status;
      this.changeStatus(this.getLeft === 0);
      this.changeTodoList(this.getTodoList);
    },
    changeTodo() {
        this.value = this.todo.name;
      if (this.todo.name.trim() == "") {
        this.delTodo(this.index);
      } else {
        this.changeTodoList(this.getTodoList);
      }
      this.editMode = false;
    },
    edit() {
      this.editMode = true;
    },
    blurEdit() {
      this.changeTodoList(this.getTodoList);
      this.editMode = false;
    },
  },
};
</script>

<style lang="scss" scoped>
.todo {
  font-size: 24px;
  padding: 0 10px;
  display: flex;
  border: 1px solid #e6e6e6;
  height: 59px;
  align-items: center;

  .todo-name {
    width: 90%;
    text-align: left;
    padding: 0 30px;
  }
  .icon {
    color: #af5b5e;
    font-size: 15px;
    cursor: pointer;
  }
}
</style>