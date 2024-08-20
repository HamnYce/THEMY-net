
//RetrieveHosts/
/*
Request: {"rows":row[]} (rows you want to create)
Response: {"rows": row[], errors: string[]} (created rows) 
*/

import { apiClient, handleResponse } from './apiClient';

interface Host {
  id: number;
  name: string;
  ip: string;
  ports: string;
}

interface FetchHostsResponse {
  hosts: Host[];
  errors?: string[];
}

export const fetchHosts = async (limit: number, offset: number): Promise<FetchHostsResponse> => {
  try {
    const response = await apiClient.post<FetchHostsResponse>('/RetrieveHosts', { limit, offset });
    return handleResponse(response);
  } catch (error) {
    console.error('Error fetching hosts:', error);
    // TODO: Improve error handling.
    throw new Error('Unable to retrieve hosts. Please try again later.');
  }
};

