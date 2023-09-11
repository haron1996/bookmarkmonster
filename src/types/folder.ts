export interface Folder {
	created_at?: string | null;
	deleted_at?: string | null;
	folder_description?: string;
	folder_id?: string;
	folder_name?: string;
	label?: string;
	path?: string;
	starred?: string;
	subfolder_of?: string | null;
	updated_at?: string;
	user_id?: string;
}
