// ── main ────────────────────────────────────────────────────────────────────⊃
// footerCollapse.js - Advanced collapse/peek behavior

document.addEventListener('DOMContentLoaded', () => {
    const footer = document.querySelector('footer');
    const toggleBtn = document.getElementById('footer-toggle');
    // Target the text span if it exists in your toggle (or adjust the ID if needed)
    const fontText = document.getElementById('font-text');

    if (toggleBtn && footer) {

        // 1. CLICK EVENT
        toggleBtn.addEventListener('click', () => {
            // If currently peeking, fully open it
            if (footer.classList.contains('peek')) {
                footer.classList.remove('peek');
                toggleBtn.innerHTML = '<span class="toparrow">⌄</span>';
                if (fontText) fontText.textContent = 'FONT';
            }
            // If collapsed, fully open it
            else if (footer.classList.contains('collapsed')) {
                footer.classList.remove('collapsed');
                toggleBtn.innerHTML = '<span class="toparrow">⌄</span>';
            }
            // Otherwise, collapse it
            else {
                footer.classList.add('collapsed');
                toggleBtn.innerHTML = '<span class="toparrow">⌃</span>';
            }
        });

        // 2. MOUSEENTER (Peek)
        toggleBtn.addEventListener('mouseenter', () => {
            // Only peek if it is currently fully collapsed
            if (footer.classList.contains('collapsed')) {
                footer.classList.replace('collapsed', 'peek');
                if (fontText) fontText.textContent = 'HUHU';
            }
        });

        // 3. MOUSELEAVE (Return to collapsed)
        toggleBtn.addEventListener('mouseleave', () => {
            if (footer.classList.contains('peek')) {
                footer.classList.replace('peek', 'collapsed');
                if (fontText) fontText.textContent = 'FONT';
            }
        });
    }
});