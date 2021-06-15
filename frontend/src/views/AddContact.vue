<template>
  <div class="background">
    <div class="header">
      <el-button
        class="back-button"
        type="primary"
        icon="el-icon-arrow-left"
        @click="$router.push(`/allcontacts`)"
        >Back</el-button
      >
      <h3 class="header-text">Add Contact</h3>
    </div>
    <div class="avatar">
      <el-avatar
        class="avatar-image"
        :size="large"
        src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
      >
      </el-avatar>
      <el-upload class="upload-demo" :limit="1" accept="image/png image/jpeg">
        <el-button class="add-avatar-button" icon="el-icon-circle-plus" circle></el-button>
      </el-upload>
    </div>
    <div class="form">
      <div class="label">Personal:</div>
      <el-input
        class="input-field"
        placeholder="Name"
        prefix-icon="el-icon-user-solid"
        v-model="contact.personal.full_name"
      >
        {{ contact.personal.full_name }}
      </el-input>
      <div class="label">Email:</div>
      <div v-for="(emailType, i) in contact.email" :key="i">
        <div v-for="(emailPhone, index) in emailType" :key="index">
          <el-input
            class="input-field"
            :placeholder="i"
            prefix-icon="el-icon-printer"
            v-model="emailType[index]"
          >
          </el-input>
        </div>
        <el-button class="save-button" type="primary" @click="AddElemToEmails(i)"
          >Add {{ i }}</el-button
        >
      </div>

      <div class="label">Phone:</div>
      <div v-for="(phoneType, i) in contact.phone" :key="i">
        <div v-for="(work, index) in phoneType" :key="index">
          <el-input
            class="input-field"
            :placeholder="i"
            prefix-icon="el-icon-s-cooperation"
            v-model="phoneType[index]"
          >
          </el-input>
        </div>
        <el-button class="save-button" type="primary" @click="AddElemToPhones(i)"
          >Add {{ i }}</el-button
        >
      </div>
    </div>
    <el-button class="save-button-main" type="primary" @click="createContact">Save</el-button>
    <br />
    <img class="footer-image" src="../assets/add-contact-image.svg" />
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "AddContact",
  data: () => ({
    contact: {
      personal: {
        full_name: ""
      },
      email: {
        work: [],
        personal: []
      },
      phone: {
        work: [],
        home: [],
        personal: []
      }
    }
  }),
  methods: {
    AddElemToPhones(group) {
      switch (group) {
        case "work":
          this.contact.phone.work.push("");
          break;
        case "home":
          this.contact.phone.home.push("");
          break;
        case "personal":
          this.contact.phone.personal.push("");
          break;
        default:
          break;
      }
    },
    AddElemToEmails(group) {
      switch (group) {
        case "work":
          this.contact.email.work.push("");
          break;
        case "personal":
          this.contact.email.personal.push("");
          break;
        default:
          break;
      }
    },
    async createContact() {
      const contactPayload = JSON.stringify(this.contact);
      try {
        const res = await axios.post(
          "/contacts/create?id=" + window.sessionStorage.getItem("sessionID"),
          JSON.parse(contactPayload)
        );
        this.$router.push("/allcontacts");
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
  background: url("../assets/background.svg") no-repeat center center fixed;
  background-size: 100%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

.label {
  font-weight: 600;
  padding-bottom: 1vh;
}

.header {
  padding-top: 3vh;
}

.back-button {
  background-color: transparent;
  color: #512da8;
  border: none;
  float: left;
}

.header-text {
  margin-top: 3vh;
  margin-right: 5vw;
}

.avatar {
  position: relative;
  margin: auto;
  /* margin-right: 5vw; */
}

.avatar-image {
  position: relative;
  margin-left: auto;
  margin-right: auto;
  margin-top: 5vh;
  margin-bottom: 5vh;
  height: 10vw;
  width: 10vw;
}

.add-avatar-button {
  position: absolute;
  bottom: 5vh;
  left: 51vw;
}

.form {
  width: 25vw;
  margin: auto;
  margin-bottom: 5vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.input-field {
  margin-bottom: 1vh;
}

.save-button {
  width: 24vw;
  height: 5vh;
  margin-bottom: 3vh;
}

.save-button-main {
  background-color: #512da8;
  border: none;
  color: white;
  width: 24vw;
  height: 5vh;
  margin-bottom: 3vh;
}

.footer-image {
  width: 0vw;
  height: 0vh;
  padding-left: auto;
  padding-right: auto;
}

@media (max-width: 1000px) and (max-height: 812px) {
  .header-text {
    margin-right: 18vw;
  }
  .avatar-image {
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5vh;
    margin-bottom: 5vh;
    height: 10em;
    width: 10em;
    align-self: center;
  }

  .add-avatar-button {
    position: absolute;
    left: 57vw;
  }

  .form {
    width: 60vw;
    margin: auto;
    margin-bottom: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .save-button {
    width: 50vw;
  }

  .footer-image {
    width: 40vw;
  }

  .background {
    background: url("../assets/background.svg") no-repeat center center fixed;
    background-size: 100%;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }
}
</style>
