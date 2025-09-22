/**
 * Validates email format
 */
export function validateEmail(email: string): boolean {
	const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	return emailRegex.test(email);
}

/**
 * Validates password strength
 */
export function validatePassword(password: string): {
	isValid: boolean;
	errors: string[];
} {
	const errors: string[] = [];
	
	if (password.length < 8) {
		errors.push('Password must be at least 8 characters long');
	}
	
	if (!/[A-Z]/.test(password)) {
		errors.push('Password must contain at least one uppercase letter');
	}
	
	if (!/[a-z]/.test(password)) {
		errors.push('Password must contain at least one lowercase letter');
	}
	
	if (!/\d/.test(password)) {
		errors.push('Password must contain at least one number');
	}
	
	return {
		isValid: errors.length === 0,
		errors
	};
}

/**
 * Validates required fields
 */
export function validateRequired(value: string, fieldName: string): string | null {
	if (!value || value.trim().length === 0) {
		return `${fieldName} is required`;
	}
	return null;
}

/**
 * Validates confirm password
 */
export function validateConfirmPassword(password: string, confirmPassword: string): string | null {
	if (password !== confirmPassword) {
		return 'Passwords do not match';
	}
	return null;
}
