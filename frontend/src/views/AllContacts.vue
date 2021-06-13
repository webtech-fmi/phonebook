<template>
  <div class='background'>
    <div class='header'>
      <el-button
        class='more-button'
        type='primary'
        icon='el-icon-more'
        @click='sideMenu = !sideMenu'
        circle
      ></el-button>
      <h2 class='UserContacts'>My Contacts</h2>
      <el-button
        class='add-button'
        type='primary'
        icon='el-icon-plus'
        circle
        @click='$router.push(`/add-contact`)'
      >
      </el-button>
    </div>
    <HamburgerMenu class='side-menu' v-if='sideMenu'></HamburgerMenu>
    <el-input class='inline-input' v-model='search' placeholder='Filter Users'>
      {{ search }}
    </el-input>
    <ul class='list'>
      <el-row class='row' v-for='contact in contacts' :key='contact'>
        <div v-if='querySearch(contact.name)'>
          <el-column class='left-column'>
            <el-button
              class='contact-button'
              @click='$router.push(`/contact/${contact.id}`)'
            >
              <h3>{{ contact.name }}</h3>
            </el-button>
          </el-column>
          <el-column class='right-column'>
            <el-button
              class='favorite-button'
              :icon='contact.favorite ? `el-icon-star-off` : `el-icon-star-on`'
              @click='toggleFavorite(contact.id)'
            >
            </el-button>
          </el-column>
        </div>
      </el-row>
    </ul>
  </div>
</template>

<script>
import HamburgerMenu from '../components/HamburgerMenu.vue';

export default {
  name: 'AllContacts',
  data: () => ({
    search: '',
    sideMenu: false,
    contacts: [
      {
        id: 1,
        name: 'Merilin Pisina',
        favorite: true,
      },
      {
        id: 2,
        name: 'Nicole Angelova',
        favorite: false,
      },
      {
        id: 3,
        name: 'Velko Bonev',
        favorite: false,
      },
      {
        id: 4,
        name: 'Kostadin Kolchev',
        favorite: false,
      },
      {
        id: 5,
        name: 'Kaloqn Bonev',
        favorite: false,
      },
    ],
  }),

  components: {
    HamburgerMenu,
  },

  methods: {
    /* eslint no-param-reassign: ['error', { 'props': false }] */
    toggleFavorite(id) {
      this.contacts.forEach((contact) => {
        if (contact.id === id) {
          contact.favorite = !contact.favorite;
        }
      });
    },
    querySearch(contact) {
      return contact.toLowerCase().indexOf(this.search.toLowerCase()) === 0;
    },
  },
};
</script>

<style scoped>
.background {
  height: 100vh;
  background-image: url('../assets/background.svg');
  background-size: 100%;
}

.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin: 3vh 3vw 3vh 3vw;
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
}
</style>
