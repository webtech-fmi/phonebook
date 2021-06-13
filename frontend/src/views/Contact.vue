<template>
  <div class='background'>
    <div class='header'>
      <el-button
        class='back-button'
        type='primary'
        icon='el-icon-arrow-left'
        size='medium'
        v-if='id != -1'
      >
        Back
      </el-button>
      <el-button
        class='more-button'
        type='primary'
        icon='el-icon-more'
        circle
        @click='sideMenu = !sideMenu'
      >
      </el-button>
      <HamburgerMenu class='side-menu' v-if='sideMenu'></HamburgerMenu>
    </div>
    <div class='avatar'>
      <el-avatar
        class='avatar-image'
        :size='large'
        src='https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
      >
      </el-avatar>
      <el-upload class='upload-demo' :limit='1' accept='image/png image/jpeg'>
        <el-button
          class='add-avatar-button'
          icon='el-icon-circle-plus'
          circle
        ></el-button>
      </el-upload>
    </div>
    <h2 class='contact-name'>{{ name }}</h2>
    <div class='info-form'>
      <div v-for='(email, index) in emails' :key='index'>
        <div class='form-field'>
          <div class='left-column'>
            <el-field class='field-header'>Email</el-field>
            <br />
            <el-field class='field-data'> {{ emails[index] }}</el-field>
            <el-divider class='divider'></el-divider>
          </div>
          <div class='right-column'>
            <el-button
              class='edit-button'
              icon='el-icon-edit'
              circle
            ></el-button>
          </div>
        </div>
      </div>
      <div v-for='(workPhone, index) in phones.workPhones' :key='index'>
        <div class='form-field'>
          <div class='left-column'>
            <el-field class='field-header'>Work Phone {{ index + 1 }}</el-field>
            <br />
            <el-field class='field-data'>
              {{ phones.workPhones[index] }}</el-field
            >
            <el-divider class='divider'></el-divider>
          </div>
          <div class='right-column'>
            <el-button
              class='edit-button'
              icon='el-icon-edit'
              circle
            ></el-button>
          </div>
        </div>
      </div>
      <div v-for='(mobilePhone, index) in phones.mobilePhones' :key='index'>
        <div class='form-field'>
          <div class='left-column'>
            <el-field class='field-header'
              >Mobile Phone {{ index + 1 }}</el-field
            >
            <br />
            <el-field class='field-data'>
              {{ phones.mobilePhones[index] }}</el-field
            >
            <el-divider class='divider'></el-divider>
          </div>
          <div class='right-column'>
            <el-button
              class='edit-button'
              icon='el-icon-edit'
              circle
            ></el-button>
          </div>
        </div>
      </div>
      <div v-for='(personalPhone, index) in phones.personalPhones' :key='index'>
        <div class='form-field'>
          <div class='left-column'>
            <el-field class='field-header'
              >Personal Phone {{ index + 1 }}</el-field
            >
            <br />
            <el-field class='field-data'>
              {{ phones.personalPhones[index] }}</el-field
            >
            <el-divider class='divider'></el-divider>
          </div>
          <div class='right-column'>
            <el-button
              class='edit-button'
              icon='el-icon-edit'
              circle
            ></el-button>
          </div>
        </div>
      </div>
    </div>
    <el-button class='merge-button' type='primary' v-if='id != -1'
      >Merge with...</el-button
    >
    <el-button class='delete-button' type='primary' v-if='id != -1'
      >Delete</el-button
    >
  </div>
</template>

<script>
import HamburgerMenu from '../components/HamburgerMenu.vue';

export default {
  name: 'Contact',
  data: () => ({
    id: -1,
    name: 'John Doe',
    emails: ['JohnDoe@gmail.com'],
    company: 'JohnDoes company',
    phones: {
      mobilePhones: ['11111111', '22222222'],
      workPhones: ['1111111111', '2222222222'],
      personalPhones: ['123123123', '123123123'],
    },
    sideMenu: false,
  }),
  components: {
    HamburgerMenu,
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

.back-button {
  background-color: transparent;
  background-color: #008080;
  border: none;
  float: left;
}

.more-button {
  background-color: #008080;
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

.edit-button {
  background-color: transparent;
  border: none;
}

.merge-button {
  background-color: #008080;
}

.delete-button {
  background-color: #008080;
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
