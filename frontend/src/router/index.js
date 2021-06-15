import { createRouter, createWebHistory } from "vue-router";
import Login from "../views/Login.vue";
import App from "../App.vue";
import SignUp from "../views/SignUp.vue";
import Contact from "../views/Contact.vue";
import AllContacts from "../views/AllContacts.vue";
import AddContact from "../views/AddContact.vue";
import Favourites from "../views/Favourites.vue";

const routes = [
  {
    path: "/",
    name: "App",
    component: App
  },
  {
    path: "/login",
    name: "Login",
    component: Login
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp
  },
  {
    path: "/me",
    name: "Me",
    component: Contact
  },
  {
    path: "/allcontacts",
    name: "AllContacts",
    component: AllContacts
  },
  {
    path: "/add-contact",
    name: "AddContact",
    component: AddContact
  },
  {
    path: "/contact",
    name: "contact",
    component: Contact,
    props: true
  },
  {
    path: "/favourite",
    name: "Favourite",
    component: Favourites,
    props: true
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
