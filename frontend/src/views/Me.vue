<template>
    <div class="background">
        <div>
            <img class="myPicture" src="../assets/person_pin-24px.svg" alt="userPicture">
            <h2 class="Someone">Someone</h2>
        </div>
        <img class="settings" src="../assets/settings-24px.svg" alt="settingPicture">
        <img class="dots" src="../assets/more_horiz-24px.svg" alt="dotsPicture">
        <div class="myInfo-form">
            <el-field class="info">Email</el-field>
            <el-input placeholder="Your email" v-model="email">{{email}}</el-input>
            <el-field class="info">Company</el-field>
            <el-input placeholder="Your company" v-model="company">{{company}}</el-input>
            <el-field class="info">Mobile</el-field>
            <el-input placeholder="Your mobile" v-model="mobile">{{mobile}}</el-input>
            <el-field class="info">Work</el-field>
            <el-input placeholder="Your work" v-model="work">{{work}}</el-input>
        </div>
    </div>
</template>

<script>
export default {
  name: 'Me',
  data: () => ({
    email: '',
    company: '',
    mobile: '',
    work: '',
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

 .myInfo-form{
     width: 50vw;
     margin:auto;
     margin-bottom: 5vh;
     padding: 1vw;
 }

 .settings{
     position: absolute;
     top: 0px;
     right: 0px;
 }

 .dots{
     position: fixed;
     top: 0px;
     left: 0px;
 }

 .myPicture{
     width: 30vw;
     height: 30vw;
     padding: 3vw;
 }

 .info{
     color: #008080;
 }

</style>
