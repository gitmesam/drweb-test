<template lang="pug">
  table.table.is-fullwidth
    thead
      tr
        th
        th File name
        th Created
    tbody(v-if="files.length")
      tr(v-for="(file, idx) in files")
        td(v-text="idx + 1")
        td
          a(:href="file.webViewLink", target="_blank") {{ file.name }}
        td(v-text="prettyDate(file.createdTime)")
    tbody(v-else)
      tr
        td(colspan="2") File list is empty
</template>

<script>
  import moment from 'moment'

  export default {
    props: {
      files: {
        type: Array,
        default: () => [],
      }
    },
    methods: {
      prettyDate (date) {
        return moment(date).fromNow();
      }
    },
  }
</script>
