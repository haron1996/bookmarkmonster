export const hideTagCreatedBookmarkForm = () => {
	const wrapper = document.getElementById('tagCreatedBookmarkWrapper') as HTMLDivElement | null;

	if (wrapper === null) return;

	wrapper.style.backgroundColor = 'rgb(0, 0, 0, 0)';

	wrapper.style.transform = 'translateX(200%)';
};
