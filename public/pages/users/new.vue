<template>
  <div>
    <b-button @click="post">POST!</b-button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data: function() {
    return {
      user: {
        name: 'nuxt'
      }
    }
  },
  methods: {
    post() {
      // axios.post("localhost:8080")
      alert('まだ未実装！')
      return
    }
  }
}
</script>
