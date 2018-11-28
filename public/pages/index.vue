<template>
  <div v-if="!loading">
    <h2>Users</h2>
    <b-table striped :items="users"></b-table>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data: function() {
    return {
      users: [],
      loading: true
    }
  },
  methods: {
    getApiData() {
      this.loading = true
      axios
        .get('http://localhost:8080/users')
        .then(response => {
          console.log(response)
          this.users = response.data.users
        })
        .catch(error => {
          console.log(error)
          this.errored = true
        })
        .finally(() => {
          this.loading = false
        })
    }
  },
  mounted() {
    this.getApiData()
  }
}
</script>

<style>
</style>
