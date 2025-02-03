import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginLayout from '@/layouts/LoginLayout.vue'
import UserSignIn from '@/views/UserSignIn.vue'
import UserSignUp from '@/views/UserSignUp.vue'
import UserSetPassword from '@/views/UserSetPassword.vue'
import { useUserStore } from '@/stores/user'
import DefaultLayout from '@/layouts/DefaultLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Bount.ing',
      component: DefaultLayout,
      children: [
        {
          path: '',
          name: 'Home',
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
          path: '/admin',
          name: 'AdminDashboard',
          component: () => import('../views/AdminDashboard.vue')
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
          path: '/profile',
          name: 'Profile',
          component: () => import('../views/UserProfileView.vue'),
          meta: { needsAuth: true }
        },
        {
          path: '/claim/:id',
          name: 'claim',
          component: () => import('../views/ClaimBountyView.vue'),
          props: true,
          meta: { needsAuth: true }
        },
        {
          path: '/pricing',
          name: 'Pricing',
          component: () => import('../views/PricingView.vue')
        },

        {
          path: '/oauth',
          name: 'OAuth',
          component: () => import('../views/AuthView.vue')
        },
        {
          path: '/connect/stripe',
          name: 'ConnectStripe',
          component: () => import('../views/ConnectStripeView.vue')
        }
      ]
    },
    {
      path: '/',
      component: LoginLayout,
      children: [
        {
          path: 'signin',
          name: 'userSignIn',
          meta: { skipIfLoggedIn: true },
          component: UserSignIn
        },
        {
          path: 'signup',
          name: 'userSignUp',
          meta: { skipIfLoggedIn: true },
          component: UserSignUp
        },
        {
          path: 'signup/verify/:mailb64/:code',
          name: 'userSetPassword',
          meta: { skipIfLoggedIn: true },
          component: UserSetPassword
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  const defaultTitle = 'Bount.ing';
  document.title = to.name ? to.name.toString() : defaultTitle;

  const userStore = useUserStore();
  
  // Handle routes that require authentication
  if (to.matched.some((record) => record.meta.needsAuth || record.meta.needsAdmin || record.meta.needsModerator)) {
    if (!userStore.isLoggedIn && to.path !== '/signup' && !to.path.startsWith('/signup/verify')) {
      next({ path: '/signin' });
      return;
    }
    next(); // Proceed if logged in
    return;
  }

  // Handle case where user is logged in and trying to access login/signup routes
  if (to.matched.some((record) => record.meta.skipIfLoggedIn) && userStore.isLoggedIn) {
    next({ path: '/profile' });
    return;
  }
  
  // Ensure the 'verify' route is treated as a special case
  if (to.path.startsWith('/signup/verify')) {
    // Check if the verification URL should be accessible (e.g., if it's just a password reset)
    if (!userStore.isLoggedIn) {
      next(); // Allow access to the verification page even if not logged in
      return;
    }
  }

  next(); // Continue to next route
});


export default router
