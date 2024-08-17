import axios from 'axios';
import { debugLog } from '@/utils/debugLogUtil';

/**
 * Fetch data from the specified JSON file path.
 * @param {string} filePath - The path to the JSON file.
 * @returns {Promise<{data: any[], error: string | null}>} - The fetched data and any error encountered.
 */
export const fetchDataFromJson = async (filePath: string): Promise<{ data: any[], error: string | null }> => {
  try {
    const response = await axios.get(filePath);
    debugLog('Fetched data:', response.data); // Log the fetched data if debug mode is true
    return { data: Array.isArray(response.data) ? response.data : [response.data], error: null };
  } catch (error) {
    debugLog('Error fetching data', error); // Log the error if debug mode is true
    if (error instanceof Error) {
      return { data: [], error: `Error fetching data: ${error.message}` };
    }
    return { data: [], error: 'Unknown error fetching data' };
  }
};
