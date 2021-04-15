<template>
<div class="background">
    <div class="header">
        <el-button
        class="back-button"
        type="primary"
        icon="el-icon-arrow-left"
        @click="$router.push('/allcontacts')">Back</el-button>
        <h3 class="header-text">Add Contact</h3>
    </div>
    <div class="avatar">
        <el-avatar class="avatar-image" :size="large" src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png">
        </el-avatar>
        <el-upload
            class="upload-demo"
            :limit="1"
            accept="image/png image/jpeg">
            <el-button class="add-avatar-button" icon="el-icon-circle-plus" circle></el-button>
        </el-upload>
    </div>
    <div class="form">
        <el-input class="input-field"
                    placeholder="Name"
                    prefix-icon="el-icon-user-solid"
                    v-model="contact.name">
        {{ contact.name }}
        </el-input>
        <el-input class="input-field"
                    placeholder="Email"
                    prefix-icon="el-icon-printer"
                    v-model="contact.email">
        {{ contact.email }}
        </el-input>
        <el-input class="input-field"
                    placeholder="Work phone"
                    prefix-icon="el-icon-s-cooperation"
                    v-model="contact.workPhone">
        {{ contact.workPhone }}
        </el-input>
        <el-input class="input-field"
                    placeholder="Mobile phone"
                    prefix-icon="el-icon-mobile-phone"
                    v-model="contact.homePhone">
        {{ contact.homePhone }}
        </el-input>
        <el-input class="input-field"
                    placeholder="Home phone"
                    prefix-icon="el-icon-phone"
                    v-model="contact.mobilePhone">
        {{ contact.mobilePhone }}
        </el-input>
    </div>
    <el-button class="save-button">Save</el-button>
    <br>
    <img class="footer-image" src="../assets/add-contact-image.svg">
</div>
</template>

<script>
export default {
  name: 'AddContact',
  data() {
    return {
      contact: {
        name: '',
        email: '',
        workPhone: '',
        homePhone: '',
        mobilePhone: '',
      },
    };
  },
};
</script>

<style scoped>
.background{
     height: 100vh;
     background-image: url('../assets/background.svg');
     background-size: 100%;
}
.header{
    margin-top: 3vh;
}

.back-button{
    background-color: transparent;
     color: #512da8;
     border: none;
     float: left;
}

.header-text{
    margin-top: 3vh;
    margin-right: 5vw;
}

.avatar{
    position: relative;
    margin: auto;
    margin-right: 5vw;
}

.avatar-image{
    position: relative;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5vh;
    margin-bottom: 5vh;
    height: 10vw;
    width: 10vw;
}

.add-avatar-button{
    position: absolute;
    bottom: 5vh;
    left: 51vw;
}

.form{
     width: 25vw;
     margin:auto;
     margin-bottom: 5vh;
     display: flex;
     flex-direction: column;
     align-items: center;
 }

 .input-field{
     margin-bottom: 1vh;
 }

  .save-button{
     width: 24vw;
     height: 5vh;
     background: #008080;
     color: white;
     margin-bottom: 3vh;
 }

 .footer-image{
     width:10vw;
     height: 20vh;
     margin-left: auto;
     margin-right: auto;
 }

 @media (max-width: 1000px) and (max-height: 812px) {
     .header-text{
         margin-right: 18vw;
     }
    .avatar-image{
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5vh;
    margin-bottom: 5vh;
    height: 10em;
    width: 10em;
    align-self: center;
    }

    .add-avatar-button{
        position: absolute;
        left: 57vw;
    }

    .form{
     width: 60vw;
     margin:auto;
     margin-bottom: 0;
     display: flex;
     flex-direction: column;
     align-items: center;
    }

    .save-button{
        width: 50vw;
    }

    .footer-image{
        width:40vw;
    }
}
</style>
