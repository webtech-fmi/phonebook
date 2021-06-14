<template>
  <div class="background">
    <div class="header">
      <el-button
        class="back-button"
        type="primary"
        icon="el-icon-arrow-left"
        size="medium"
        v-if="id != -1"
        @click="$router.go(-1)"
      >
        Back
      </el-button>
      <el-button
        class="more-button"
        type="primary"
        icon="el-icon-more"
        circle
        @click="sideMenu = !sideMenu"
      >
      </el-button>
      <HamburgerMenu class="side-menu" v-if="sideMenu"></HamburgerMenu>
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
    <h2 class="contact-name">{{ user.personal.full_name }}</h2>
    <div class="info-form">
      <div v-for="(emailType, i) in user.email" :key="i">
        <div v-for="(email, index) in emailType" :key="index">
          <div class="form-field">
            <div class="left-column">
              <div class="field-header">{{ i }} email {{ index + 1 }}</div>
              <br />
              <el-input
                class="field-data"
                v-model="emailType[index]"
                :disabled="editButton"
                >{{
              }}</el-input>
              <el-divider class="divider"></el-divider>
            </div>
            <div class="right-column"></div>
          </div>
        </div>
      </div>

      <div v-for="(phoneType, i) in user.phone" :key="i">
        <div v-for="(phone, index) in phoneType" :key="index">
          <div class="form-field">
            <div class="left-column">
              <el-field class="field-header"> {{ i }} phone {{ index + 1 }}</el-field>
              <br />
              <el-input
                class="field-data"
                v-model="phoneType[index]"
                :disabled="editButton"
                >{{
              }}</el-input>
              <el-divider class="divider"></el-divider>
            </div>
            <div class="right-column"></div>
          </div>
        </div>
      </div>
    </div>

    <el-button class="edit-button" type="primary" @click="editButton = !editButton">
      {{ editButton == true ? "Edit" : "Save" }}</el-button
    >
    <el-button class="merge-button" type="primary" v-if="id != -1">Merge with...</el-button>
    <el-button class="delete-button" type="primary" v-if="id != -1">Delete</el-button>
  </div>
</template>

<script>
import HamburgerMenu from "../components/HamburgerMenu.vue";
import axios from "axios";

export default {
  name: "Contact",
  props: ["id"],
  data: () => ({
    user: {
      id: "",
      user_id: "",
      email: [],
      phone: [],
      personal: {
        full_name: ""
      }
    },
    sideMenu: false,
    editButton: true
  }),
  components: {
    HamburgerMenu
  },
  methods: {
    async getProfile() {
      const session = `{
            "session_id": "${window.sessionStorage.getItem("sessionID")}"
        }`;
      try {
        const res = await axios.post("/profiles/get", JSON.parse(session));
        this.user = res.data;
      } catch (e) {
        console.warn(e);
      }
    },
    async getContact() {
      try {
        const res = await axios.get("/contacts/by-id?id=" + this.id);
        this.user = res.data;
      } catch (e) {
        console.warn(e);
      }
    }
  },
  async mounted() {
    if (!!this.id) {
      await this.getContact();
    } else {
      await this.getProfile();
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

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin: 3vh 3vw 3vh 3vw;
}

.back-button {
  border: none;
  float: left;
}

.avatar {
  margin-bottom: 5vh;
}
.avatar-image {
  position: relative;
  margin-left: auto;
  margin-right: auto;
  height: 10vw;
  width: 10vw;
}

.contact-name {
  margin-bottom: 3vh;
}

.info-form {
  width: 50vw;
  margin: auto;
  margin-bottom: 5vh;
  padding: 1vw;
}

.form-field {
  display: flex;
  flex-direction: row;
  margin-bottom: 2vh;
}

.right-column {
  width: 10%;
}

.left-column {
  width: 90%;
  text-align: left;
}

.field-header {
  opacity: 0.65;
}

.divider {
  margin: 1vw 0 1vw 0;
}

@media (max-width: 1000px) and (max-height: 812px) {
  .avatar-image {
    display: block;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5vh;
    height: 10em;
    width: 10em;
    align-self: center;
  }
  .аvatar {
    margin-bottom: 5vh;
  }

  .info-form {
    width: 70vw;
    margin: auto;
    margin-bottom: 5vh;
    padding: 1vw;
  }
}
</style>
