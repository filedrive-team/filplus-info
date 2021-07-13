import axios from 'axios'
const instance = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL
})
instance.interceptors.request.use((config) => {
  return config
})
instance.interceptors.response.use(response => {
  const res = response.data
  return res
})
export default function fetch(options) {
  const { url, method = 'post', data, params } = options
  return new Promise((resolve, reject) => {
    if (method === 'get') {
      instance.get(url, {
        params: params
      }).then(res => resolve(res)).catch(e => reject(e))
    } else {
      instance.post(url, data).then(res => resolve(res)).catch(e => reject(e))
    }
  })
}
