<template>
  <div class="background">
    <div>
      <img class="avatar" src="../assets/avatar.png" alt="Avatar" />
    </div>
    <h2 class="phone-book-header">PhoneBook</h2>
    <h3 class="sign-in-header">Sign In</h3>
    <br />
    <h4 class="sign-in-message">Hi there! nice to see you again</h4>
    <div class="login-form">
      <el-input class="input-field" placeholder="Please input email" v-model="email">
        {{ email }}
      </el-input>
      <br />
      <el-input
        class="input-field"
        placeholder="Please input password"
        v-model="password"
        show-password
      >
        {{ password }}
      </el-input>
    </div>
    <div>
      <el-button class="sign-in-button" @click="login">Sign in</el-button>
    </div>

    <div class="footer">
      <el-button class="forgot-password">Forgot password?</el-button>
      <el-button class="sign-up-button" @click="$router.push('/signup')">Sign up</el-button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data: () => ({
    email: "",
    password: ""
  }),
  methods: {
    async login() {
      const user = `{
      "email": "${this.email}",
      "password":"${this.password}"
      }`;
      try {
        const res = await axios.post("/auth/login", JSON.parse(user));
        if (res.status == 200) {
            window.sessionStorage.setItem("sessionID", res.data.id);
            this.$router.push('/me');
        } else {
            console.log("TODO");
        }
      } catch (e) {
          console.warn(e);
      }
    }
  }
};
</script>

<style scoped>
.background {
  height: 100vh;
  background-image: url("../assets/background.svg");
  background-size: 100%;
}
.avatar {
  height: 20vh;
  padding: 1em;
}

.phone-book-header {
  color: #512da8;
  padding-bottom: 2em;
}

.sign-in-header {
  color: #008080;
}

.sign-in-message {
  color: #b3b3ba;
  margin-bottom: 5vh;
}
.login-form {
  width: 25vw;
  margin: auto;
  margin-bottom: 5vh;
}
.sign-in-button {
  width: 24vw;
  height: 5vh;
  background: #008080;
  color: white;
}
.footer {
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 7vh;
}
@media (max-width: 800px) and (max-height: 600px) {
  .login-form {
    width: 70vw;
    margin: auto;
    margin-bottom: 3vh;
  }
  .sign-in-button {
    width: 40vw;
  }
}
.input-field {
  margin-bottom: 3vh;
}
</style>
