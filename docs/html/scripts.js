function toggleSidebar() {
  document.getElementsByTagName("body")[0].classList.toggle("sidebar-collapsed");
}

function adjustSidebar() {
  if (window.innerWidth < 760) {
    document.getElementsByTagName("body")[0].classList.add("sidebar-collapsed");
  }
}

window.addEventListener("load", adjustSidebar);
window.addEventListener("resize", adjustSidebar);
