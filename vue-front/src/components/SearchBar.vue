<template lang="pug">
  div
    .field
      label Write the filter query
    form(action="#", @submit.prevent)
      .field.has-addons
        .control
          input.input(type="text", v-model="search", placeholder="filter")
        .control
          input.button.is-info(type="submit", ref="btn", @click="fetchFiles(search)", value="Search")
      p.has-text-danger(v-if="error") {{ error }}
      slot
</template>

<script>
  export default {
    props: {
      fetchFiles: {
        type: Function,
        default: () => {
          alert('not implemented')
        }
      },
      error: {
        type: String,
        default: () => null,
      }
    },
    data: () => ({
      search: '',
    }),
  }
</script>
