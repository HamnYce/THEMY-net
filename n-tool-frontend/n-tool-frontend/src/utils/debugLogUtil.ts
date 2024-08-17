/**
 * Log messages to the console if debug mode is enabled.
 * @param {string} message - The message to log.
 * @param {any} [data] - Optional data to log with the message.
 */
export const debugLog = (message: string, data?: any) => {
    if (process.env.NEXT_PUBLIC_DEBUG_MODE === 'true') {
      if (data !== undefined) {
        console.log(message, data);
      } else {
        console.log(message);
      }
    }
  };
  