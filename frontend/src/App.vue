<template>
  <div :class="theme" class="bg-ctp-base text-ctp-text min-h-screen font-inter">
    <the-header></the-header>
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component class="flex justify-center" :is="Component" />
      </keep-alive>
    </router-view>
    <!-- <the-footer></the-footer> -->
  </div>
</template>

<script>
import { RouterView } from 'vue-router'
import TheHeader from './components/TheHeader.vue'
import TheFooter from './components/TheFooter.vue'

export default {
  components: {
    TheHeader,
    TheFooter,
    RouterView
  },
  data() {
    return {
      theme: 'ctp-frappe'
    }
  }
}
</script>
