import axios from 'axios'

const baseURL = axios.create({
    api: "localhost:5000/api/v1/"
})

export default baseURL