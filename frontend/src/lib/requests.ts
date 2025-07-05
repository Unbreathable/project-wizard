import { getServerUrl } from "$lib";

/**
 * Make a POST request to the specified URL with a JSON body
 * @param path The path to send the request to
 * @param body The request body (will be JSON stringified)
 * @returns Promise resolving to a map of the response data
 */
export async function postRequestURL(path: string, body: any): Promise<Record<string, any>> {
	try {
		// Send the request
		const response = await fetch(getServerUrl() + path, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(body)
		});

		// Check if the response is ok
		if (!response.ok) {
			return {
				success: false,
				message: `HTTP error! status: ${response.status}`,
			}
		}

		// Parse and return the JSON response
		const data: Record<string, any> = await response.json();
		return data;
	} catch (error) {
		// Re-throw the error to be handled by the caller
		return {
			success: false,
			message: `Failed to send request: ${error}`
		}
	}
}