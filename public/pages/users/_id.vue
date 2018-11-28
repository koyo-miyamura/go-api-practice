<template>
  <b-form class="form" @submit.prevent="postUser">
      <h2>Edit User</h2>
      <b-form-group label="Name:">
        <b-form-input type="text"
                      v-model="user.name"
                      required
                      placeholder="Enter name">
        </b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Submit</b-button>
    </b-form>
</template>

<style scoped>
.form {
  margin-bottom: 30px;
}
</style>

<script>
import axios from 'axios'
axios.defaults.headers.common['Content-Type'] = `application/json`
export default {
  data: function() {
    return {
      user: {},
      url: 'http://localhost:8080/users/' + this.$route.params.id
    }
  },
  methods: {
    getApiData() {
      axios
        .get(this.url)
        .then(response => {
          this.user = response.data.user
        })
        .catch(error => {
          console.log(error)
        })
    },
    updateUser() {
      axios
        .put(this.url, {
          name: this.form.name
        })
        .then(response => {
          console.log('Updated' + response.data)
        })
        .catch(error => {
          console.log(error)
        })
        .finally(() => {
          this.getApiData()
        })
    }
  },
  mounted() {
    this.getApiData()
  }
}
</script>
