<template>
  <the-title title="Home Page"></the-title>
  <section>
    <div>
      <ul class="flex justify-between">
        <li
          class="shadow-lg rounded-lg p-4 bg-ctp-mantle m-4"
          :key="pageEntry.id"
          v-for="pageEntry in pageEntries"
        >
          <h4 class="text-xl text-ctp-green">{{ pageEntry.title }}</h4>
          <p>
            <a class="underline text-ctp-lavender" :href="pageEntry.link">{{ pageEntry.link }}</a>
          </p>
        </li>
      </ul>
    </div>
  </section>
</template>

<script>
import TheTitle from '@/components/TheTitle.vue'
export default {
  components: {
    TheTitle
  },
  data() {
    return {
      pageEntries: [
        {
          id: '1234',
          title: 'test',
          link: 'http://vedaant.dev'
        },
        {
          id: '24847',
          title: 'Hello World',
          link: 'http://vedantwnkhade.atlassian.dev'
        },
        {
          id: '1wry04i4',
          title: 'Google',
          link: 'http://googl.ecom'
        }
      ]
    }
  }
}
</script>
