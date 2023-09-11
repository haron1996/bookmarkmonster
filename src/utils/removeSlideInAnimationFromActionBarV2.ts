export function removeSlideInAnimationFromActionBarV2() {
	const div = document.getElementById('actionBar') as HTMLDivElement | null;

	if (div === null) return;

	div.classList.remove('animate__animated');

	div.classList.remove('animate__fadeInDown');
}
