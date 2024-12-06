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
      path: "/profile",
      name: "Profile",
      component: () => import('../views/UserProfileView.vue'),
      meta: { needsAuth: true }
    },
    {
      path: "/pricing",
      name: "Pricing",
      component: () => import('../views/PricingView.vue')
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
    },
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

    const userStore = useUserStore()
    if (
      to.matched.some(
        (record) => record.meta.needsAuth || record.meta.needsAdmin || record.meta.needsModerator
      )
    ) {
      // this route requires auth, check if logged in
      // if not, redirect to signin page
      if (!userStore.isLoggedIn) {
        next({ path: '/signin' })
      
        return
      } else {
        next() // go to wherever I'm going
        return
      }
    } else if (to.matched.some((record) => record.meta.needsAdmin)) {
      //Check if user is admin here, else redirect to home
    } else if (to.matched.some((record) => record.meta.needsModerator)) {
      //Check if user is moderator here, else redirect to home
    } else {
      if (to.matched.some((record) => record.meta.skipIfLoggedIn) && userStore.isLoggedIn) {
        //in case user is logged in, redirect to dashboard instead of showing this component
        next({ path: '/user' })
        return
      }
      next() // does not require auth, make sure to always call next()!
    }
  })
  

export default router
