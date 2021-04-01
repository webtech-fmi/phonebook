import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import Home from '../views/Home.vue';
import SignUp from '../views/SignUp.vue';
import Contact from '../views/Contact.vue';
import AllContacts from '../views/AllContacts.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp,
  },
  {
    path: '/me',
    name: 'Me',
    component: Contact,
  },
  {
    path: '/allcontacts',
    name: 'AllContacts',
    component: AllContacts,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
