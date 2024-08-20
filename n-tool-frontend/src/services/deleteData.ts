
// DeleteHosts/
/*
Request: {"rowIDs": number[]} (row ids to delete)
Response: {"deletedRowIDs": number[], errors: string[]}
*/

import { apiClient, handleResponse } from './apiClient';

export const deleteHosts = async (hostIDs: number[]) => {
  try {
    const response = await apiClient.post('/DeleteHosts', { hostIDs });
    return handleResponse(response);
  } catch (error) {
    console.error('Error deleting hosts:', error);
    throw error;
  }
};
