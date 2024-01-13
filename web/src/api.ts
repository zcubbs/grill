import axios, {AxiosResponse} from 'axios';

export const API_BASE_URL = import.meta.env.VITE_API_URL;


export const pingApi = async (): Promise<string> => {
  try {
    const response: AxiosResponse<string> = await axios.get(`${API_BASE_URL}/v1/ping`);
    return response.data;
  } catch (error) {
    console.error('Error fetching ping:', error);
    throw error;
  }
}
