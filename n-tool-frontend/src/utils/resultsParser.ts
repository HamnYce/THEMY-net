// utils/resultsParser.ts

/**
 * Format cell value for display.
 * @param {any} value - The cell value to format.
 * @returns {string} - The formatted cell value.
 */
export const formatCellValue = (value: any): string => {
    if (Array.isArray(value)) {
      return value.map(item => (typeof item === 'object' ? JSON.stringify(item) : item)).join(', ');
    }
    if (typeof value === 'object' && value !== null) {
      return JSON.stringify(value);
    }
    return value;
  };
  