export const hideTagCreatedAlert = () => {
	const tagCreatedAlert = document.getElementById('tagCreated') as HTMLDivElement | null;

	if (tagCreatedAlert === null) return;

	tagCreatedAlert.style.transform = 'translateY(-100%)';

	tagCreatedAlert.style.opacity = '0';
};
