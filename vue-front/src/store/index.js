import Vue from 'vue'
import Vuex from 'vuex'
import { t } from './types'
import api from '../api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    files: [],
    loading: false,
    error: null,
  },
  mutations: {
    [t.FILES_LOADING] (state) {
      state.loading = true
      state.error = null
    },
    [t.FILES_ERROR] (state, error) {
      state.loading = false
      state.error = error
    },
    [t.FILES_LOADED] (state, files) {
      state.loading = false
      state.error = null
      Vue.set(state, 'files', files)
    },
  },
  actions: {
    fetchFiles({ commit }, query) {
      commit(t.FILES_LOADING)
      return api.drive.files(query)
        .then(({ data }) => {
          commit(t.FILES_LOADED, data.files || [])
        })
        .catch(({ response }) => {
          if (response && 'data' in response) {
            const { error } = response.data
            return commit(t.FILES_ERROR, error || response.data)
          }
          commit(t.FILES_ERROR, 'Fatal error occured')
        })
    }
  }
})
