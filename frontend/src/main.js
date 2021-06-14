import { createApp } from "vue";
import ElementPlus from "element-plus";
import "element-plus/lib/theme-chalk/index.css";
import App from "./App.vue";
import router from "./router";
import axios from "axios";
import VueAxios from "vue-axios";

axios.defaults.baseURL = "http://localhost:3000"

createApp(App)
  .use(VueAxios, axios)
  .use(ElementPlus)
  .use(router)
  .mount("#app");
