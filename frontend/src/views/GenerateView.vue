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
          <label for="title">Page Title</label>
          <input type="text" name="title" id="title" v-model="title" />
        </div>
        <div>
          <label for="parentId">Parent Page ID</label>
          <input type="text" name="parentId" id="parentId" v-model="parentId" />
        </div>
        <div>
          <label for="spaceId">Space ID</label>
          <input type="text" name="spaceId" id="spaceId" v-model="spaceId" />
        </div>
        <div>
          <button type="submit">Generate</button>
        </div>
      </form>
    </div>
  </section>
  <section>
    <a v-if="pageLink != ''" :href="pageLink">{{ pageLink }}</a>
  </section>
</template>

<script>
export default {
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
