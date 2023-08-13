export const showCreateTagComponent = () => {
	const createTagComponent = document.getElementById('createTag') as HTMLDivElement | null;

	if (createTagComponent === null) return;

	const overlay = document.getElementById('overlay') as HTMLDivElement | null;

	if (overlay === null) return;

	overlay.style.transform = 'translateX(0)';

	overlay.style.opacity = '1';

	createTagComponent.style.transform = 'translateX(0)';
};
