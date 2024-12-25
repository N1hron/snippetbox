const navLinks = document.querySelectorAll("nav a");

for (let link of navLinks) {
	if (link.href == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}