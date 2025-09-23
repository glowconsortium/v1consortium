/**
 * Formats a date for display
 */
export function formatDate(dateString: string, locale: string = 'en-US'): string {
	try {
		const date = new Date(dateString);
		return new Intl.DateTimeFormat(locale, {
			year: 'numeric',
			month: 'short',
			day: 'numeric',
			hour: '2-digit',
			minute: '2-digit'
		}).format(date);
	} catch {
		return 'Invalid date';
	}
}

/**
 * Formats a relative time (e.g., "2 hours ago")
 */
export function formatRelativeTime(dateString: string, locale: string = 'en-US'): string {
	try {
		const date = new Date(dateString);
		const now = new Date();
		const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000);
		
		if (diffInSeconds < 60) {
			return 'Just now';
		} else if (diffInSeconds < 3600) {
			const minutes = Math.floor(diffInSeconds / 60);
			return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
		} else if (diffInSeconds < 86400) {
			const hours = Math.floor(diffInSeconds / 3600);
			return `${hours} hour${hours > 1 ? 's' : ''} ago`;
		} else if (diffInSeconds < 2592000) {
			const days = Math.floor(diffInSeconds / 86400);
			return `${days} day${days > 1 ? 's' : ''} ago`;
		} else {
			return formatDate(dateString, locale);
		}
	} catch {
		return 'Invalid date';
	}
}

/**
 * Formats file size in human readable format
 */
export function formatFileSize(bytes: number): string {
	if (bytes === 0) return '0 Bytes';
	
	const k = 1024;
	const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
	const i = Math.floor(Math.log(bytes) / Math.log(k));
	
	return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

/**
 * Formats number with thousands separator
 */
export function formatNumber(num: number, locale: string = 'en-US'): string {
	return new Intl.NumberFormat(locale).format(num);
}

/**
 * Truncates text to specified length
 */
export function truncateText(text: string, maxLength: number): string {
	if (text.length <= maxLength) return text;
	return text.slice(0, maxLength - 3) + '...';
}
