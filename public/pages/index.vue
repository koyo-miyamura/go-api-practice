<template>
  <div v-if="!loading">
    <h2>Users</h2>
    <b-alert variant="danger"
             dismissible
             :show="isError"
             @dismissed="isError=false">
      Error status {{ error.response.status }} {{ error.response.statusText }}
    </b-alert>
    <b-table striped :items="users" :fields="fields">
      <template slot="name" slot-scope="row">
        <nuxt-link :to="{ path: 'users/' + row.item.id }">{{row.item.name}}</nuxt-link>
      </template>
      <template slot="action" slot-scope="row">
        <b-button variant="danger" size="sm" @click.stop="deleteUser(row.item)">
          Delete
        </b-button>
      </template>
    </b-table>
    <b-form class="form" @submit.prevent="postUser">
      <h2>New User</h2>
      <b-form-group label="Name:">
        <b-form-input type="text"
                      v-model="form.name"
                      required
                      placeholder="Enter name">
        </b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Submit</b-button>
    </b-form>
  </div>
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
      users: [],
      fields: ['id', 'name', 'created_at', 'updated_at', 'action'],
      loading: true,
      form: {
        name: ''
      },
      url: 'http://localhost:8080/users',
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
      this.loading = true
      axios
        .get(this.url)
        .then(response => {
          console.log(response)
          this.users = response.data.users
        })
        .catch(error => {
          console.log(error)
        })
        .finally(() => {
          this.loading = false
        })
    },
    deleteUser(user) {
      if (!confirm(`delete ${user.name}?`)) {
        return
      }
      this.loading = true
      this.isError = false
      axios
        .delete(`${this.url}/${user.id}`)
        .then(response => {
          console.log(response)
        })
        .catch(error => {
          console.log(error)
          this.isError = true
        })
        .finally(() => {
          this.loading = false
          this.getApiData()
        })
    },
    postUser() {
      this.isError = false
      axios
        .post(this.url, {
          name: this.form.name
        })
        .then(response => {
          console.log('Created' + response.data)
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

<style>
</style>
