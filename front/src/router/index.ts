import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Bount.ing',
      component: HomeView
    },
    {
      path: '/about',
      name: 'About Bount.ing',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/contact',
      name: 'Contact Bount.ing',
      component: () => import('../views/ContactView.vue')
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('../views/DashboardView.vue')
    },
    {
      path: '/privacy',
      name: 'Privacy Policy',
      component: () => import('../views/PrivacyView.vue')
    },
    {
      path: '/terms',
      name: 'Terms of Use',
      component: () => import('../views/TermsView.vue')
    },
    {
      path: '/help',
      name: 'Help & F.A.Q.',
      component: () => import('../views/HelpFAQView.vue')
    },
    {
      path: "/profile",
      name: "Profile",
      component: () => import('../views/UserProfileView.vue')
    },
    {
      path: "/login",
      name: "Login",
      component: () => import('../views/LoginView.vue')
    },
    {
      path: "/auth",
      name: "Auth",
      component: () => import('../views/AuthView.vue')
    },
    {
      path: "/connect/stripe",
      name: "ConnectStripe",
      component: () => import('../views/ConnectStripeView.vue')
    }
  ]
})

router.beforeEach((to, from, next) => {
  const defaultTitle = 'Bount.ing';
  document.title = to.name ? to.name.toString() : defaultTitle;
  next();
});
export default router
