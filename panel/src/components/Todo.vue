<template>
  <div id="todo-container">
    <router-view style="box-shadow:none"/>
    <div id="footer">
      <span>{{getLeft}} item left</span>
      <a-radio-group default-value="all" v-model="tag" @change="tagChange" id="button-group">
        <a-radio-button value="all">All</a-radio-button>
        <a-radio-button value="active">Active</a-radio-button>
        <a-radio-button value="completed">Completed</a-radio-button>
      </a-radio-group>
      <span id="clear" @click="delComplete" v-if="getTodoList.length>getLeft">Clear completed</span>
    </div>
  </div>
</template>

<script>
import { mapGetters,mapActions } from "vuex";
export default {
  name: "Todos",
  data() {
    return {
      tag: "all",
    };
  },
  computed: {
    ...mapGetters(["getTodoList", "getLeft"]),
  },
  methods:{
      ...mapActions(['changeTodoList']),
      tagChange(e){
          this.tag = e.target.value;
          this.$router.push({path:'/'+this.tag});
      },
      delComplete(){
          this.changeTodoList(this.getTodoList.filter(todo=>todo.status==false))
      }
  }
};
</script>

<style lang="scss" scoped>
#todo-container {
  background: white;
  #footer {
    height: 50px;
    display: flex;
    align-items: center;
    padding: 0 15px;
    #button-group{
        margin-left:70px;
        width:200px;
        display: flex;
        justify-content: space-between;
    }
    #clear{
        margin-left: 50px;
        cursor: pointer;
    }
    #clear:hover{
        text-decoration: underline;
    }
  }
}
</style>