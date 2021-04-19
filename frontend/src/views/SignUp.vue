<template>
    <div class="background">
        <h2 class="sign-up-header">Sign up</h2>
        <div class="sign-up-form">
            <el-field class="info" style="color:#512da8">Name</el-field>
            <el-input placeholder="Your name" v-model="name">{{name}}</el-input>
            <el-field class="info" style="color:#512da8">Phone</el-field>
            <el-input placeholder="Your phone number" v-model="number">{{number}}</el-input>
            <el-field class="info" style="color:#512da8">Email</el-field>
            <el-input placeholder="Your email address" v-model="email">{{email}}</el-input>
            <el-field class="info" style="color:#512da8">Password</el-field>
            <el-input placeholder="Your password" v-model="password">{{password}}</el-input>
        </div>
        <div>
            <el-checkbox v-model="privacyPolicyCheck">
                I agree to the <el-link href="" target="#512da8">Terms of services</el-link> and
                <el-link href="">Privacy Policy</el-link>
            </el-checkbox>
        </div>
        <div>
            <el-button class="continue-button">Continue</el-button>
        </div>
        <el-field type="question" >Have an account?</el-field>
        <el-link @click="$router.push('/login')" target="_self">Sign In</el-link>
        <div>
            <img class="SignUpPicture" src="../assets/sign-up-image.svg" alt="Picture">
        </div>
    </div>
</template>

<script>
export default {
  name: 'SignUp',
  data: () => ({
    privacyPolicyCheck: false,
    name: '',
    number: '',
    email: '',
    password: '',
  }),
};
</script>

<style scoped>
 .background{
     height: 100vh;
     background-image: url('../assets/background.svg');
     background-size: 100%;
 }

 .sign-up-header{
     margin-top: 5vh;
     margin-bottom: 3vh;
     color: #008080;
 }

 .sign-up-form{
     width: 30vw;
     margin:auto;
     margin-bottom: 5vh;
     padding: 1vw;
 }

 .continue-button{
     width: 24vw;
     height: 5vh;
     background: #008080;
     color: white;
     margin: 3vw;
 }

 .SignUpPicture {
     height: 20vh;
     padding: 1em;
 }

 a:any-link{
     color: #512da8;
     font-weight: bold;
 }
 .info{
     font-weight: bold;
 }

</style>
