
/*
Purpose: Centralize the logic for making HTTP requests. 
This file can handle setting up the base URL, 
configuring default headers, managing authentication tokens, 
and handling common error responses.

Usage: All CRUD operations will use the 
functions defined in this file to make HTTP requests,
ensuring consistent API interactions.

*/

import axios, { AxiosInstance, AxiosResponse } from 'axios';

const apiClient: AxiosInstance = axios.create({
  baseURL: 'http://localhost:8080', 
  headers: {
    'Content-Type': 'application/json',
  },
});



//Test function
const handleResponse = <T>(response: AxiosResponse<T>): T => response.data;

export { apiClient, handleResponse };
