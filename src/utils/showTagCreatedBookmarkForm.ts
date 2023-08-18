export const showTagCreatedBookmarkForm = () => {
	const wrapper = document.getElementById('tagCreatedBookmarkWrapper') as HTMLDivElement | null;

	if (wrapper === null) return;

	wrapper.style.transform = 'translateX(0)';

	wrapper.style.backgroundColor = 'rgb(0, 0, 0, .6)';
};
