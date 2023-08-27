export function closeTagMenu() {
	const menu = document.getElementById('tagMenu') as HTMLElement | null;

	if (menu === null) return;

	menu.style.display = 'none';
}
