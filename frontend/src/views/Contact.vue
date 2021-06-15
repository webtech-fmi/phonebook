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
      <el-button
        class="more-button"
        type="primary"
        icon="el-icon-more"
        circle
        @click="sideMenu = !sideMenu"
      >
      </el-button>
      <HamburgerMenu class="side-menu" @clicked="OnMenuClose" v-if="sideMenu"></HamburgerMenu>
    </div>
    <div class="avatar">
      <el-avatar
        class="avatar-image"
        :size="large"
        src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
      >
      </el-avatar>
    </div>
    <div>
      <span class="contact-name">{{ user.personal.full_name }}</span>
    </div>
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
      <div class="form-field" v-if="user.metadata.organization !== undefined">
        <div class="left-column">
          <el-field class="field-header"> Organization </el-field>
          <br />
          <el-input
            class="field-data"
            v-model="user.metadata.organization"
            :disabled="editButton"
            >{{
          }}</el-input>
          <el-divider class="divider"></el-divider>
        </div>
      </div>

      <div class="form-field" v-if="user.metadata.address !== undefined">
        <div class="left-column">
          <el-field class="field-header"> Address </el-field>
          <br />
          <el-input
            class="field-data"
            v-model="user.metadata.address"
            :disabled="editButton"
            >{{
          }}</el-input>
          <el-divider class="divider"></el-divider>
        </div>
      </div>
    </div>

    <div class="dropdown-button">
      <el-dropdown>
        <el-button @click="editButton = !editButton" type="primary">
          Add Additional Info<i class="el-icon--right"></i>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="addAddress">Add Address</el-dropdown-item>
            <el-dropdown-item @click="addOrganization">Add Organization</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <br />

    <el-button class="footer-button" type="primary" @click="edit">
      {{ editButton == true ? "Edit" : "Save" }}</el-button
    >
    <el-button class="footer-button" @click="openMerge" type="primary" v-if="isContact"
      >Merge with...</el-button
    >
    <el-dialog
      class="merge-dialog"
      title="Merge with..."
      v-model="mergeDialog"
      :before-close="handleClose"
    >
      <el-checkbox v-model="contact.checked" v-for="contact in contacts" :key="contact.id">{{
        contact.personal.full_name
      }}</el-checkbox>
      <el-button type="primary" @click="mergeContacts">Confirm</el-button>
    </el-dialog>
    <el-button class="footer-button" type="primary" v-if="isContact" @click="deleteContact"
      >Delete</el-button
    >
    <el-button
      class="footer-button"
      type="primary"
      v-if="isContact && user.favourite == false"
      @click="favourite"
      >Add To Favourites</el-button
    >
    <el-button
      class="footer-button"
      type="primary"
      v-if="isContact && user.favourite == true"
      @click="favourite"
      >Remove From Favourites</el-button
    >
  </div>
</template>

<script>
import HamburgerMenu from "../components/HamburgerMenu.vue";
import axios from "axios";

export default {
  name: "Contact",
  props: ["id"],
  data: () => ({
    mergeDialog: false,
    contacts: [],
    user: {
      id: "",
      user_id: "",
      email: [],
      phone: [],
      personal: {
        full_name: ""
      },
      metadata: {
        address: null,
        organization: null
      }
    },
    sideMenu: false,
    editButton: true,
    isContact: false
  }),
  components: {
    HamburgerMenu
  },
  methods: {
    async favourite() {
      this.user.favourite = !this.user.favourite;
      const favouritePayload = {
        session_id: window.sessionStorage.getItem("sessionID"),
        id: this.id,
        favourite: this.user.favourite
      };
      try {
        const res = await axios.post("/contacts/favourite", favouritePayload);
        if (res.status == 200) {
          this.$router.push({ name: `contact`, params: { id: this.id } });
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }
    },
    async deleteContact() {
      const deletePayload = {
        session_id: window.sessionStorage.getItem("sessionID"),
        id: this.id,
      };
      try {
        const res = await axios.post("/contacts/delete", deletePayload);
        if (res.status == 200) {
          this.$router.push('/allcontacts');
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }
    },
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
    async getContacts() {
      try {
        const res = await axios.get(
          "/contacts/by-owner?id=" + window.sessionStorage.getItem("sessionID")
        );
        this.contacts = res.data.contacts;
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
    },
    addOrganization() {
      this.user.metadata.organization = "";
    },
    addAddress() {
      this.user.metadata.address = "";
    },
    async editContact() {
      const editPayload = {
        session_id: window.sessionStorage.getItem("sessionID"),
        id: this.id,
        email: this.user.email,
        phone: this.user.phone,
        personal: this.user.personal,
        metadata: this.user.metadata
      };

      try {
        const res = await axios.post("/contacts/edit", editPayload);
        if (res.status == 200) {
          this.$router.push({ name: `contact`, params: { id: this.id } });
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }
    },
    async editProfile() {
      const editPayload = {
        session_id: window.sessionStorage.getItem("sessionID"),
        email: this.user.email,
        phone: this.user.phone,
        personal: this.user.personal,
        metadata: this.user.metadata
      };

      try {
        const res = await axios.post("/profiles/edit", editPayload);
        if (res.status == 200) {
          this.$router.push("/me");
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }
    },
    async edit() {
      this.editButton = !this.editButton;
      if (this.editButton) {
        if (this.id) {
          await this.editContact();
        } else {
          await this.editProfile();
        }
      }
    },
    async mergeContacts() {
      const mergePayload = {
        session_id: window.sessionStorage.getItem("sessionID"),
        contacts: this.contacts.filter(c => c.checked).map(c => c.id),
        main: this.id
      };

      try {
        const res = await axios.post("/contacts/merge", mergePayload);
        if (res.status == 200) {
          this.$router.push("/allcontacts");
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }

      this.mergeDialog = false;
    },
    async openMerge() {
      await this.getContacts();
      this.contacts.map(element => {
        element.checked = false;
        return element;
      });
      console.info(this.contacts);
      this.mergeDialog = true;
    },
    OnMenuClose() {
      this.sideMenu = false;
    }
  },
  async mounted() {
    if (!!this.id) {
      await this.getContact();
      this.isContact = true;
    } else {
      await this.getProfile();
    }
  }
};
</script>

<style scoped>
.favorite-button {
  position: relative;
  background-color: transparent;
  border: none;
  z-index: 1;
  font-size: 20px;
}

.background {
  /* height: 100vh; */
  background: url("../assets/background.svg") no-repeat center center fixed;
  background-size: 100%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 3vh 3vw 3vh 3vw;
}

.dropdown-button {
  margin-top: 1vh;
  margin-bottom: 1vh;
}

.el-checkbox:last-of-type {
  margin-right: 30px;
}

.back-button {
  background-color: transparent;
  color: #512da8;
  border: none;
  float: left;
}

.footer-button {
  margin: 1vh 1vw 5vh 1vw;
  width: 20vh;
}

.avatar {
  margin-bottom: 2vh;
}

.avatar-image {
  position: relative;
  margin-left: auto;
  margin-right: auto;
  height: 10vw;
  width: 10vw;
}

.contact-name {
  margin-bottom: 4vh;
  font-size: 30px;
  /* margin-left: 5vh; */
}

.info-form {
  width: 50vw;
  margin: auto;
  /* margin-bottom: 5vh; */
  padding: 1vw;
}

.form-field {
  display: flex;
  flex-direction: row;
  margin-bottom: 2vh;
}

.right-column {
  width: 0%;
}

.left-column {
  width: 100%;
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
    width: 75vw;
    margin: auto;
    /* margin-bottom: 5vh; */
    /* padding: 1vw; */
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
