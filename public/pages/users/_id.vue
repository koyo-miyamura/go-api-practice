<template>
  <b-form class="form" @submit.prevent="updateUser">
    <h2>Edit User</h2>
    <b-alert variant="danger"
             dismissible
             :show="isError"
             @dismissed="isError=false">
      Error status {{ error.response.status }} {{ error.response.statusText }}
    </b-alert>
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
      url: 'http://localhost:8080/users/' + this.$route.params.id,
      error: {
        response: {
          status: '',
          statusText: ''
        }
      },
      isError: false
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
      this.isError = false
      axios
        .put(this.url, {
          name: this.user.name
        })
        .then(response => {
          console.log('Updated' + response.data)
          this.$router.push('/')
        })
        .catch(error => {
          console.log(error)
          this.error = error
          this.isError = true
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
