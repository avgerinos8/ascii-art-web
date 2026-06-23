// LocalStorage key identifier for state persistence
const STORAGE_KEY = 'ascii_web_settings';

// Reads all current UI values and commits them to localStorage
function saveAllToStorage() {
    const state = {};

    const activeFont = document.querySelector('input[name="font"]:checked');
    if (activeFont) state.font = activeFont.value;

    const wrapToggle = document.getElementById('font-wrap-toggle');
    if (wrapToggle) state.fontWrap = wrapToggle.checked;

    const alignSelect = document.getElementById('font-align-select');
    if (alignSelect) state.fontAlign = alignSelect.value;

    const realtimeToggle = document.getElementById('realtime-toggle');
    if (realtimeToggle) state.realtime = realtimeToggle.checked;

    // Extract background slider layout matrix values dynamically
    ['BgSize', 'BgDistance', 'BreatheSpeed', 'Noise', 'BreatheDepth'].forEach(name => {
        const el = document.querySelector(`input[name="${name}"]`);
        if (el) state[name] = el.value;
    });

    localStorage.setItem(STORAGE_KEY, JSON.stringify(state));
}

// Intercepts DOM values and injects stored properties, checking for manual reloads
function restoreAllFromStorage() {
    // Check if the page load was triggered by a manual refresh (F5, Ctrl+F5, Reload button)
    const navigationEntry = performance.getEntriesByType('navigation')[0];
    if (navigationEntry && navigationEntry.type === 'reload') {
        localStorage.removeItem(STORAGE_KEY);
        return; // Halt restoration to allow the page to reset to its native defaults
    }

    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) return;

    try {
        const state = JSON.parse(raw);

        if (state.font) {
            const radio = document.querySelector(`input[name="font"][value="${state.font}"]`);
            if (radio) radio.checked = true;
        }

        const wrapToggle = document.getElementById('font-wrap-toggle');
        if (wrapToggle && state.fontWrap !== undefined) wrapToggle.checked = state.fontWrap;

        const alignSelect = document.getElementById('font-align-select');
        if (alignSelect && state.fontAlign) alignSelect.value = state.fontAlign;

        const realtimeToggle = document.getElementById('realtime-toggle');
        if (realtimeToggle && state.realtime !== undefined) realtimeToggle.checked = state.realtime;

        ['BgSize', 'BgDistance', 'BreatheSpeed', 'Noise', 'BreatheDepth'].forEach(name => {
            const el = document.querySelector(`input[name="${name}"]`);
            if (el && state[name] !== undefined) el.value = state[name];
        });
    } catch (e) {
        console.error("Error parsing storage targets:", e);
    }
}

// Execute injection synchronously right away so elements hold correct values on load
restoreAllFromStorage();

// Hook listeners once DOM evaluation completes
document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('footer form');
    if (form) {
        // Save state whenever any option or text input transitions layout metrics
        form.addEventListener('change', saveAllToStorage);
        form.addEventListener('submit', saveAllToStorage);
    }

    // Clean slate boundary validation mapping for the Reset button
    const resetBtn = document.getElementById('reset-btn');
    if (resetBtn) {
        resetBtn.addEventListener('click', () => {
            localStorage.removeItem(STORAGE_KEY);
        });
    }
});
