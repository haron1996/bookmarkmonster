export const showDuplicateTagAlert = () => {
	const tagCreatedAlert = document.getElementById('duplicateTag') as HTMLDivElement | null;

	if (tagCreatedAlert === null) return;

	tagCreatedAlert.style.transform = 'translateY(0)';

	tagCreatedAlert.style.opacity = '1';
};
