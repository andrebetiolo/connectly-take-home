import 'https://cdn.jsdelivr.net/npm/axios@1.2.2/dist/axios.min.js';

const baseService = axios.create({
  baseURL: API_URL
});

baseService.defaults.timeout = 60000;

baseService.interceptors.request.use(config => {
  return config
}, error => {
  return Promise.reject(error)
})

baseService.interceptors.response.use(response => {
  let message = response.data && response.data.message;
  if (message) showMessage(message);
  return response;
}, error => {
  let message = 'Ocorreu um erro ao fazer o procedimento ou tem algum problema de conexão com o servidor';
  let status;
  let data;
  if (error.response && error.response.data) data = error.response.data;

  if (data && data.message) message = data.message;
  if (error.response && error.response.status) status = error.response.status;

  if (status == 401) {
    clearToken();
  }

  const ID_ERRO = 'errors';
  if (data && data[ID_ERRO]) {
    message = 'Existe(m) erro(s) no(s) parâmetro(s) abaixo: \n\n';
    const fieldWithError = Object.keys(data[ID_ERRO]);
    fieldWithError.forEach(key => {
      message += key + ': ' + data[ID_ERRO][key] + '\n\n';
    })
  }

  showMessage(message, 'error');
  return Promise.reject(error);
})

const http = {
  get(uri, params = {}) {
    return baseService.get(uri, { params });
  },

  post(uri, params) {
    return baseService.post(uri, params);
  },

  put(uri, params) {
    return baseService.put(uri, params);
  },

  delete(uri) {
    return baseService.delete(uri);
  }
}

window.http = http;

export default http;
