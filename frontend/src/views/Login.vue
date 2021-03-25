<template>
<div class="background">
    <div>
      <img class="avatar" src="../assets/avatar.png" alt="Avatar">
    </div>
    <h2 class="phone-book-header">PhoneBook</h2>
    <h3 class="sign-in-header">Sign In</h3>
    <br>
        <h4 class="sign-in-message">Hi there! nice to see you again</h4>
    <div class = "login-form">
        <el-input placeholder="Please input email" v-model="email">{{ email }}</el-input>
        <br>
        <el-input placeholder="Please input password" v-model="password" show-password>
            {{ password }}
        </el-input>
    </div>
    <div>
        <el-button class="sign-in-button">Sign in</el-button>
    </div>

    <div class="footer">
        <el-button class="forgot-password">Forgot password?</el-button>
        <el-button class="sign-up-button">Sign up</el-button>

    </div>
</div>
</template>

<script>
export default {
  name: 'Login',
  data: () => ({
    email: '',
    password: '',
  }),
};
</script>

<style scoped>
 .background{
     height: 100vh;
     background-image: url('../assets/background.svg');
     background-repeat: no-repeat;
     background-size: 100%;
 }
 .avatar {
     height: 20vh;
     padding: 1em;
 }

 .phone-book-header{
     color: #512da8;
     padding-bottom: 2em;
 }

 .sign-in-header{
     color: #008080;
 }

 .sign-in-message{
     color: #b3b3ba;
     margin-bottom: 5vh;
 }
 .login-form{
     width: 30vw;
     margin:auto;
     margin-bottom: 5vh;
 }
 .sign-in-button{
     width: 24vw;
     height: 5vh;
     background: #008080;
     color: white;
 }
 .footer{
     display: flex;
     flex-direction: row;
     justify-content: center;
     margin-top: 10vh;
 }
</style>
