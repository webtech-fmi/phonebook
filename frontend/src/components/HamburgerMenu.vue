<template>
  <div class="container">
    <el-button class="close-button" icon="el-icon-close" size="medium"></el-button>

    <div class="user-info">
      <el-avatar
        class="avatar-image"
        :size="large"
        src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
      ></el-avatar>
      <h4 class="user-name">{{ name }}</h4>
      <h5 class="user-email">{{ email }}</h5>
    </div>

    <el-menu
      default-active="2"
      class="menu"
      @open="handleOpen"
      @close="handleClose"
      :collapse="isCollapse"
    >
      <el-menu-item index="1" @click="$router.push('/me')">
        <template #title>
          <i class="el-icon-house"></i>
          <span>Me</span>
        </template>
      </el-menu-item>
      <el-menu-item index="2" @click="$router.push('/allcontacts')">
        <i class="el-icon-s-custom"></i>
        <template #title>Contacts</template>
      </el-menu-item>
      <el-menu-item index="3">
        <i class="el-icon-star-on"></i>
        <template #title>Favorites</template>
      </el-menu-item>
      <el-menu-item index="4" @click="logout">
        <i class="el-icon-switch-button"></i>
        <template #title>Sign Out</template>
      </el-menu-item>
    </el-menu>

    <img class="footer-image" src="../assets/hamburger-menu-image.svg" alt="FooterImg" />
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HamburgerMenu",
  data: () => ({
    name: "John Doe",
    email: "JohnDoe@gmail.com"
  }),
  methods: {
    async logout() {
        const session = `{
            "session_id": "${window.sessionStorage.getItem("sessionID")}"
        }`;
      try {
        const res = await axios.post("/auth/logout", JSON.parse(session));
        if (res.status == 200) {
          window.sessionStorage.clear();
          this.$router.push("/login");
        } else {
          console.log("TODO");
        }
      } catch (e) {
        console.warn(e);
      }
    }
  }
};
</script>

<style scoped>
.container {
  position: absolute;
  display: flex;
  flex-direction: column;
  justify-content: center;
  z-index: 999;
  left: 0;
  top: 0;
  width: 40vw;
  height: 100vh;
  background-color: #fff;
}

.container > * {
  margin-bottom: 4vh;
}

.menu > * {
  margin-bottom: 4vh;
}

.close-button {
  width: 30vw;
  background-color: transparent;
  color: black;
  border: none;
  position: absolute;
  top: 3vh;
  left: 0;
}

.avatar-image {
  width: 10vw;
  height: 10vw;
}

@media (max-width: 1024px) and (max-height: 824px) {
  .container {
    width: 60vw;
  }

  .avatar-image {
    width: 30vw;
    height: 30vw;
  }
}
</style>
