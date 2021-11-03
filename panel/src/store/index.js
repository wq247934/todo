import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    todoList: [],
    status:false
  },
  mutations: {
    addTodoList(state, todos) {
      state.todoList.push(todos);
    },
    delTodo(state,index){
        state.todoList.splice(index,1);
    },
    changeTodoList(state,todoList){
        state.todoList=todoList;
    },
    changeStatus(state,status){
        state.status = status;
    }
  },
  actions: {
    addTodoList(context, todos) {
      context.commit("addTodoList", todos);
      window.localStorage.setItem('todos',JSON.stringify(this.state.todoList));
    },
    delTodo(context,index){
        context.commit("delTodo",index);
        window.localStorage.setItem('todos',JSON.stringify(this.state.todoList));
    },
    changeTodoList(context,todoList){
        context.commit("changeTodoList",todoList);
        window.localStorage.setItem('todos',JSON.stringify(this.state.todoList))
    },
    changeStatus(context,status){
        context.commit('changeStatus',status);
    }
  },
  getters:{
    getTodoList(state){
        return state.todoList;
    },
    getLeft(state){
        return state.todoList.filter(todo=>todo.status==false).length;
    },
    getStatus(state){
        return state.status;
    }
  },
  modules: {},
});
