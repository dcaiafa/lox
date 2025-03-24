function toggleSidebar() {
  document.getElementById("container").classList.toggle("container-sidebar-collapsed");
}

function adjustSidebar() {
  console.log("adjustSidebar " + String(window.innerWidth));
	if (window.innerWidth < 760) {
  	document.getElementById("container").classList.add("container-sidebar-collapsed");	
  }
}

window.addEventListener("load", adjustSidebar);
window.addEventListener("resize", adjustSidebar);
