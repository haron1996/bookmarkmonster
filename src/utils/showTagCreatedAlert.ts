export const showTagCreatedAlert = () => {
	const tagCreatedAlert = document.getElementById('tagCreated') as HTMLDivElement | null;

	if (tagCreatedAlert === null) return;

	tagCreatedAlert.style.transform = 'translateY(0)';

	tagCreatedAlert.style.opacity = '1';
};
