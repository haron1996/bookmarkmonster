export const hideAddBookmarkComponent = () => {
	const addBookmarkComponent = document.getElementById('addBookmark') as HTMLDivElement | null;

	if (addBookmarkComponent === null) return;

	addBookmarkComponent.style.transform = 'translateX(200%)';

	const overlay = document.getElementById('overlay') as HTMLDivElement | null;

	if (overlay === null) return;

	overlay.style.transform = 'translateX(-200%)';

	overlay.style.opacity = '0';
};
