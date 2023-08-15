export interface Bookmark {
	added?: Date;
	bookmark?: string;
	deleted?: Date | null;
	favicon?: string;
	host?: string;
	id?: string;
	notes?: string | null;
	thumbnail?: string;
	title?: string;
	updated?: Date | null;
	user_id?: string;
}
