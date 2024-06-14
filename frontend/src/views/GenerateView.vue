<template>
  <section>
    <div>
      <h1>Generate TSR</h1>
    </div>
    <div>
      <form @submit.prevent="submit">
        <div>
          <label for="atlassianInstance">Atlassian Instance Server</label>
          <input
            type="text"
            name="atlassianInstance"
            id="atlassianInstance"
            v-model="atlassianInstance"
          />
        </div>
        <div>
          <label for="atlassianUsername">Atlassian Username</label>
          <input
            type="text"
            name="atlassianUsername"
            id="atlassianUsername"
            v-model="atlassianUsername"
          />
        </div>
        <div>
          <label for="atlassianToken">Atlassian Token</label>
          <textarea name="atlassianToken" id="atlassianToken" v-model="atlassianToken" />
        </div>
        <div>
          <label for="jql">JQL</label>
          <input type="text" name="jql" id="jql" v-model="jql" />
        </div>
        <div>
          <button type="submit">Generate</button>
        </div>
      </form>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      atlassianInstance: '',
      atlassianUsername: '',
      atlassianToken: '',
      jql: ''
    }
  },
  methods: {
    submit() {
      console.log(this.atlassianInstance, this.atlassianUsername, this.atlassianToken, this.jql)
    }
  }
}
</script>
