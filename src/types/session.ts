export interface Session {
	AccessToken: string;
	RefreshToken: string;
	User: {
		account_password: string | null;
		created_at: Date;
		deleted: Date | null;
		email: string;
		email_verified: boolean;
		id: string;
		last_login: Date;
		name: string;
		picture: string;
		refres_token: string;
	};
}
