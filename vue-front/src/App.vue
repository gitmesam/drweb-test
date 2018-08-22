<template lang="pug">
  #app
    section.hero.is-info
      .hero-body
        h1.title.is-3 Google Drive
        p Don't forget to configure service before
        p Set the proper 'serviceBaseUrl' in src/api/index.js
        p
          small Не стал делать авторизацию в гугл-драйве через этот фронт, т.к. не нашел этого в требованиях
    section.section
      .field
        search-bar(:fetchFiles="fetchFiles", :error="error")
          pulse-loader(:loading="loading")
      files(:files="files")
</template>

<script>
import { mapState, mapActions } from 'vuex'
import PulseLoader from 'vue-spinner/src/PulseLoader.vue'
import SearchBar from './components/SearchBar.vue'
import Files from './components/Files.vue'

export default {
  name: 'app',
  computed: {
    ...mapState({
      files: state => state.files,
      loading: state => state.loading,
      error: state => state.error,
    })
  },
  methods: {
    ...mapActions({
      fetchFiles: 'fetchFiles',
    })
  },
  components: {
    PulseLoader,
    SearchBar,
    Files,
  }
}
</script>

<style>
  #app {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
</style>
