import axios from 'axios'

const serviceBaseUrl = 'http://localhost:8081'

export default {
  drive: {
    files (filter = null) {
      return axios.post(`${serviceBaseUrl}/files`, { filter })
    }
  }
}