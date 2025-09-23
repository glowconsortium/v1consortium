// API response types
// export interface ApiResponse<T = unknown> {
// 	code: number;
// 	message: string;
// 	data?: T;
// }

// export interface ApiError {
// 	code: number;
// 	message: string;
// 	details?: string;
// 	field_errors?: Record<string, string[]>;
// }

export interface PaginatedResponse<T> {
	items: T[];
	total: number;
	page: number;
	per_page: number;
	total_pages: number;
}

// export interface UsageMetrics {
// 	submissions_count: number;
// 	endpoints_count: number;
// 	storage_used: number;
// 	current_period_start: string;
// 	current_period_end: string;
// 	plan_limits: {
// 		max_submissions: number;
// 		max_endpoints: number;
// 		max_storage: number;
// 	};
// }

// export interface SubscriptionPlan {
// 	id: string;
// 	name: string;
// 	description: string;
// 	price: number;
// 	currency: string;
// 	interval: 'month' | 'year';
// 	features: string[];
// 	limits: {
// 		submissions_per_month: number;
// 		endpoints: number;
// 		storage_gb: number;
// 		team_members: number;
// 		api_keys: number;
// 		data_retention_days: number;
// 	};
// }

// export interface Subscription {
// 	id: string;
// 	organization_id: string;
// 	plan: SubscriptionPlan;
// 	status: 'active' | 'canceled' | 'past_due' | 'trialing';
// 	current_period_start: string;
// 	current_period_end: string;
// 	trial_end?: string;
// 	cancel_at?: string;
// 	created_at: string;
// }
