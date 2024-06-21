import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const login = (credentials) => {
  return axios.post(`${API_BASE_URL}/login`, credentials);
};

export const register = (credentials) => {
  return axios.post(`${API_BASE_URL}/register`, credentials);
};

export const getProfile = (token) => {
  return axios.get(`${API_BASE_URL}/profile`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
