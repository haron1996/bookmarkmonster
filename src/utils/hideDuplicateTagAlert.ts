export const hideDuplicateTagAlert = () => {
	const tagCreatedAlert = document.getElementById('duplicateTag') as HTMLDivElement | null;

	if (tagCreatedAlert === null) return;

	tagCreatedAlert.style.transform = 'translateY(-100%)';

	tagCreatedAlert.style.opacity = '0';
};
