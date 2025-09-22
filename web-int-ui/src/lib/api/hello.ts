import { apiClient } from './client.js';

export class HelloService {
	async hello() {
		return await apiClient.hello.getApiPublicV1Hello();
	}
}

// Create a singleton instance
export const helloService = new HelloService();
