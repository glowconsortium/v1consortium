
export 	interface SidebarItem {
		label: string;
		href: string;
		icon: string;
		active?: boolean;
		badge?: string;
	}


export interface HeaderUser {
    name: string;
    email: string;
    avatarUrl?: string;
}