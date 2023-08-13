export const hideOverlay = () => {
	const overlay = document.getElementById('overlay') as HTMLDivElement | null;

	if (overlay === null) return;

	const createTagComponent = document.getElementById('createTag') as HTMLDivElement | null;

	if (createTagComponent === null) return;

	const createBookmarkComponent = document.getElementById('addBookmark') as HTMLDivElement | null;

	if (createBookmarkComponent === null) return;

	createTagComponent.style.transform = 'translateX(-200%)';

	createBookmarkComponent.style.transform = 'translateX(200%)';

	overlay.style.transform = 'translateX(-200%)';

	overlay.style.opacity = '0';
};
