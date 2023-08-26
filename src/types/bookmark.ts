export interface Bookmark {
	added?: Date;
	bookmark?: string;
	deleted?: string | null;
	favicon?: string;
	host?: string;
	id?: string;
	notes?: string | null;
	thumbnail?: string;
	title?: string;
	updated?: Date;
	user_id?: string;
}
