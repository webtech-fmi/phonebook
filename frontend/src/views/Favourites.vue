<template>
  <div class="background">
    <div class="header">
      <el-button
        class="more-button"
        type="primary"
        icon="el-icon-more"
        @click="sideMenu = !sideMenu"
        circle
      ></el-button>
      <h2 class="UserContacts">Favourites</h2>
      <el-button
        class="add-button"
        type="primary"
        icon="el-icon-plus"
        circle
        @click="$router.push('/add-contact')"
      >
      </el-button>
    </div>
    <HamburgerMenu class="side-menu" @clicked="OnMenuClose" v-if="sideMenu"></HamburgerMenu>
    <el-autocomplete
      class="inline-input"
      v-model="state1"
      :fetch-suggestions="querySearch"
      placeholder="Please Input"
      @select="handleSelect"
    >
    </el-autocomplete>
    <ol class="list">
      <el-row class="row" v-for="contact in contacts" :key="contact">
        <el-column class="left-column">
          <el-button
            class="contact-button"
            @click="$router.push({ name: `contact`, params: { id: contact.id } })"
          >
            <h3>{{ contact.personal.full_name }}</h3>
          </el-button>
        </el-column>
        <el-column class="right-column">
          <el-button
            class="favorite-button"
            :icon="contact.favorite ? 'el-icon-star-off' : 'el-icon-star-on'"
          >
          </el-button>
        </el-column>
      </el-row>
    </ol>
  </div>
</template>

<script>
import HamburgerMenu from "../components/HamburgerMenu.vue";
import axios from "axios";

export default {
  name: "AllContacts",
  data: () => ({
    search: "",
    sideMenu: false,
    contacts: []
  }),

  async mounted() {
    await this.getContacts();
  },

  components: {
    HamburgerMenu
  },

  methods: {
    /* eslint no-param-reassign: ["error", { "props": false }] */
    toggleFavorite(id) {
      this.contacts.forEach(contact => {
        if (contact.id === id) {
          contact.favorite = !contact.favorite;
        }
      });
    },
    async getContacts() {
      try {
        const res = await axios.get(
          "/contacts/favourites?id=" + window.sessionStorage.getItem("sessionID")
        );
        this.contacts = res.data.contacts;
      } catch (e) {
        console.warn(e);
      }
    },
    OnMenuClose() {
      this.sideMenu = false;
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

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 3vh 3vw 3vh 3vw;
}

.more-button {
  left: 0;
}

.add-button {
  right: 0;
}

.inline-input {
  width: 40vw;
}

.row {
  margin: auto;
  width: 40vw;
  margin-bottom: 3vh;
}

.right-column {
  width: 10%;
}

.left-column {
  width: 90%;
  text-align: left;
}

.favorite-button {
  position: relative;
  background-color: transparent;
  border: none;
  z-index: 1;
}

.contact-button {
  background-color: transparent;
  border: none;
  max-width: 40vw;
  color: black;
  z-index: 0;
}

.contact-name {
  overflow: hidden;
}

.list {
  margin-top: 3vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

@media (max-width: 1024px) and (max-height: 824px) {
  .row {
    width: 70vw;
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
