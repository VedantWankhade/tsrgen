<template>
  <the-title title="Generate TSR"></the-title>
  <section class="flex justify-center">
    <div class="max-w-2xl p-6 bg-ctp-mantle shadow-md rounded-lg w-full ml-4">
      <form @submit.prevent="submit" class="flex flex-col">
        <div class="mt-2">
          <label for="atlassianInstance">Atlassian Instance Server</label>
          <input
            required
            type="text"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
            name="atlassianInstance"
            id="atlassianInstance"
            v-model="atlassianInstance"
          />
        </div>
        <div class="mt-2">
          <label for="atlassianUsername">Atlassian Username</label>
          <input
            required
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
            type="text"
            name="atlassianUsername"
            id="atlassianUsername"
            v-model="atlassianUsername"
          />
        </div>
        <div class="mt-2">
          <label for="atlassianToken">Atlassian Token</label>
          <textarea
            required
            name="atlassianToken"
            id="atlassianToken"
            class="mt-1 h-28 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
            v-model="atlassianToken"
          />
        </div>
        <div class="mt-2">
          <label for="jql">JQL</label>
          <input
            required
            type="text"
            name="jql"
            id="jql"
            v-model="jql"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
          />
        </div>
        <div class="mt-2">
          <label for="title">Page Title</label>
          <input
            required
            type="text"
            name="title"
            id="title"
            v-model="title"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
          />
        </div>
        <div class="mt-2">
          <label for="parentId">Parent Page ID</label>
          <input
            type="text"
            name="parentId"
            id="parentId"
            v-model="parentId"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
          />
        </div>
        <div class="mt-2">
          <label for="spaceId">Space ID</label>
          <input
            required
            type="text"
            name="spaceId"
            id="spaceId"
            v-model="spaceId"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-ctp-red border-ctp-mantle focus:border-ctp-red sm:text-sm bg-ctp-base"
          />
        </div>
        <div class="mt-2">
          <button
            class="w-full inline-flex justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-gradient-to-r from-ctp-red to-ctp-yellow hover:from-ctp-red hover:to-ctp-rosewater focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-ctp-blue"
            type="submit"
          >
            Generate
          </button>
        </div>
      </form>
    </div>
    <div
      v-if="pageLink != ''"
      class="text-ctp-green max-w-2xl mx-auto h-fit ml-4 mr-4 p-6 bg-ctp-mantle shadow-md rounded-lg w-full"
    >
      <a :href="pageLink">{{ pageLink }}</a>
    </div>
  </section>
</template>

<script>
import TheTitle from '@/components/TheTitle.vue'
export default {
  mounted() {
    this.atlassianToken = localStorage.getItem('token')
    this.atlassianUsername = localStorage.getItem('username')
    this.atlassianInstance = localStorage.getItem('instance')
    this.jql = localStorage.getItem('jql')
    this.parentId = localStorage.getItem('parentId')
    this.spaceId = localStorage.getItem('spaceId')
  },
  components: {
    TheTitle
  },
  data() {
    return {
      atlassianInstance: '',
      atlassianUsername: '',
      atlassianToken: '',
      jql: '',
      title: '',
      parentId: '',
      spaceId: '',
      pageLink: ''
    }
  },
  methods: {
    submit() {
      console.log(
        this.atlassianInstance,
        this.atlassianUsername,
        this.atlassianToken,
        this.jql,
        this.title,
        this.spaceId,
        this.parentId
      )
      fetch('http://localhost:8080/generate', {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          atlassianToken: this.atlassianToken,
          atlassianUsername: this.atlassianUsername,
          atlassianInstance: this.atlassianInstance,
          jql: this.jql,
          title: this.title,
          spaceId: this.spaceId,
          parentId: this.parentId
        })
      })
        .then((res) => {
          if (res.ok) {
            localStorage.setItem('token', this.atlassianToken)
            localStorage.setItem('username', this.atlassianUsername)
            localStorage.setItem('instance', this.atlassianInstance)
            localStorage.setItem('jql', this.jql)
            localStorage.setItem('parentId', this.parentId)
            localStorage.setItem('spaceId', this.spaceId)
            return res.json()
          } else {
            throw new Error('err')
          }
        })
        .then((data) => (this.pageLink = data))
        .catch((err) => console.log(err))
    }
  }
}
</script>
