
//UpdateHosts/
/*
Request: {"rows": {id: number, ...any}} (rows with id required and any other attributes to update)
Response: {"rows": row[], errors: string[]} (created rows)
*/

import { apiClient, handleResponse } from './apiClient';

interface Host {
  id: number;
  
}

export const updateHosts = async (hosts: Host[]) => {
  try {
    const response = await apiClient.post('/UpdateHosts', { hosts });
    return handleResponse(response);
  } catch (error) {
    console.error('Error updating hosts:', error);
    throw error;
  }
};
