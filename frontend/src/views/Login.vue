<template>
    <md-avatar>
      <img class="account-box-icon-3" src="../assets/avatar.png" alt="Avatar">
    </md-avatar>
    <div class="heading-4">PhoneBook</div>
    <div class="heading-1">Sign In</div>
    <br>
        <span>Hi there! nice to see you again</span>
    <div id = "login-form">
    </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      email: String,
    };
  },
};
</script>

<style scoped>
.rectangle-8-copy-3 {
width: 280px;
height: 48px;
padding: 8px 8px 8px 8px;
box-shadow: 0px 5px 10px rgba(81, 45, 168, 0.23);
background: #008080;
}
.label-1-copy-3 {
width: 280px;
height: 40px;
color: #ffffff;
font-family: "Asap";
font-weight: 600;
font-size: 17px;
text-align: center;
}
.heading-1 {
height: 36px;
color: #008080;
font-family: "Quicksand";
font-weight: 700;
font-size: 26px;
}
.heading-2 {
width: 220px;
height: 30px;
color: #989eb1;
font-family: "Asap";
font-weight: 400;
font-size: 16px;
text-align: left;
}
.rectangle-49 {
width: 280px;
height: 2px;
padding: 8px 8px 8px 8px;
background: #ebebeb;
}
.label-34 {
width: 36px;
height: 19px;
color: #512da8;
font-family: "Asap";
font-weight: 600;
font-size: 14px;
text-align: left;
}
.label-39 {
width: 150px;
height: 22px;
color: #35424a;
font-family: "Asap";
font-weight: 500;
font-size: 16px;
text-align: left;
}
.rectangle-50 {
width: 280px;
height: 2px;
padding: 8px 8px 8px 8px;
background: #ebebeb;
}
.label-35 {
width: 58px;
height: 19px;
color: #512da8;
font-family: "Asap";
font-weight: 600;
font-size: 14px;
text-align: left;
}
.label-40 {
width: 133px;
height: 33px;
color: #35424a;
font-family: "Asap";
font-weight: 500;
font-size: 24px;
letter-spacing: 4px;
text-align: left;
}
.label-1-copy-5 {
width: 132px;
height: 35px;
color: #512da8;
font-family: "Asap";
font-weight: 600;
font-size: 16px;
text-align: center;
}
.heading-4 {
height: 33px;
margin-bottom: 15px;
color: #512da8;
font-family: "Asap";
font-weight: 600;
font-size: 24px;
letter-spacing: 1px;
text-align: center;
}
.label-1-copy-9 {
width: 132px;
height: 35px;
color: #989eb1;
font-family: "Asap";
font-weight: 500;
font-size: 16px;
text-align: center;
}
.remove-red-eye-icon-1 {
width: 24px;
height: 16px;
background: #d1d6db;
}
.account-box-icon-3 {
width: 129px;
height: 129px;
}
</style>
