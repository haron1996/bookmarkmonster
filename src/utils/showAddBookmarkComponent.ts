export const showCreateBookmarkComponent = () => {
	const overlay = document.getElementById('overlay') as HTMLDivElement | null;

	if (overlay === null) return;

	const createBookmarkComponent = document.getElementById('addBookmark') as HTMLDivElement | null;

	if (createBookmarkComponent === null) return;

	overlay.style.transform = 'translateX(0)';

	overlay.style.opacity = '1';

	createBookmarkComponent.style.transform = 'translateX(0)';
};
