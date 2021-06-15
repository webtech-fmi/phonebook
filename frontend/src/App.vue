<template>
  <div>
    <!--<router-link to ="/login">Login</router-link>-->
    <Login v-if="!loggedIn"> </Login>
    <router-view v-if="loggedIn" />
  </div>
</template>

<script>
import Login from "./views/Login";

export default {
  components: {
    Login
  },
  computed: {
    loggedIn() {
      console.log(!!window.sessionStorage.getItem("sessionID"))
      return !!window.sessionStorage.getItem("sessionID");
    }
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
