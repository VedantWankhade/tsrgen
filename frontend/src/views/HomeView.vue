<template>
  <the-title title="Home Page"></the-title>
  <section>
    <div>
      <ul class="flex justify-between flex-wrap">
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
      pageEntries: []
    }
  },
  mounted() {
    fetch('http://localhost:8080/entries', {
      headers: {
        Accept: 'application/json'
      },
      method: 'GET'
    })
      .then((res) => {
        if (res.ok) {
          return res.json()
        } else {
          throw new Error('error fetching db entries')
        }
      })
      .then((data) => {
        data.forEach((entry) =>
          this.pageEntries.push({
            id: entry.pageId,
            title: entry.pageTitle,
            link: entry.pageLink
          })
        )
      })
      .catch((err) => console.log(err))
  }
}
</script>
