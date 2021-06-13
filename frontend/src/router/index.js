import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import SignUp from '../views/SignUp.vue';
import Contact from '../views/Contact.vue';
import AllContacts from '../views/AllContacts.vue';
import AddContact from '../views/AddContact.vue';

const routes = [
  // {
  //   path: '/',
  //   name: 'Home',
  //   component: Home,
  // },
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
  {
    path: '/add-contact',
    name: 'AddContact',
    component: AddContact,
  },
  {
    path: '/contact/:id',
    name: 'Contact',
    component: Contact,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
