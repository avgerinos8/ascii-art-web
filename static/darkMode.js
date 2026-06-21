// Helper function to manage local dark theme state and storage
function toggleTheme(isDark) {
    if (isDark) {
        document.body.setAttribute("data-theme", "dark");
        localStorage.setItem("theme", "dark");
    } else {
        document.body.removeAttribute("data-theme");
        localStorage.setItem("theme", "light");
    }

    const el = document.getElementById('theme-toggle');
    if (el) {
        el.checked = isDark;
    }

    // Keep the global STATE object in sync for traditional form submissions if it exists
    if (typeof STATE !== 'undefined') {
        STATE.is_dark = isDark;
    }
}

// Initial theme bootstrap from localStorage executed immediately to prevent flashing
toggleTheme(localStorage.getItem("theme") === "dark");

// Bind event tracking once the DOM layout components are fully ready
document.addEventListener("DOMContentLoaded", () => {
    const themeToggle = document.getElementById('theme-toggle');

    if (themeToggle) {
        themeToggle.addEventListener('change', (e) => {
            toggleTheme(e.target.checked);
        });
    }
});
