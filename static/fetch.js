// Compact Global State with ideal default values configured strictly (Realtime defaults to false)
// ── main ────────────────────────────────────────────────────────────────────⊃
const STATE = { font_wrap: 'on', font_align: 'left', active_font: 'standard', realtime: false, user_text: '', max_chars: 50 };

// Sync all DOM elements into STATE object directly using standard if statements
function syncState() {
    const wrapToggle = document.getElementById('font-wrap-toggle');
    const alignSelect = document.getElementById('font-align-select');
    const checkedFont = document.querySelector('input[name="font"]:checked');
    const realtimeToggle = document.querySelector('input[name="Realtime"]') || document.getElementById('realtime-toggle');
    const textarea = document.getElementById('user-text');
    const submitBtn = document.querySelector('.ascButton-submit') || document.querySelector('button[type="submit"]');
    const maxCharsInput = document.getElementById('max-chars-input');

    if (wrapToggle) { STATE.font_wrap = wrapToggle.checked ? 'on' : ''; }

    if (alignSelect) {
        STATE.font_align = alignSelect.value || 'left';
    } else {
        STATE.font_align = 'left';
    }

    if (checkedFont) {
        STATE.active_font = checkedFont.value || 'standard';
    } else {
        STATE.active_font = 'standard';
    }

    if (realtimeToggle) { STATE.realtime = realtimeToggle.checked; }
    if (textarea) { STATE.user_text = textarea.value; }
    if (maxCharsInput) { STATE.max_chars = parseInt(maxCharsInput.value) || 50; }
    if (submitBtn && realtimeToggle) { submitBtn.disabled = STATE.realtime; }
}

// Compact Fetch Request for typography settings and text input
function sendUpdate() {
    syncState();

    // Strict block condition: If realtime is disabled, halt all background fetches immediately
    if (!STATE.realtime) { return; }

    const { font_wrap, font_align, active_font, realtime, user_text, max_chars } = STATE;
    fetch("/api/session-state", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ font_wrap, font_align, active_font, realtime, user_text, max_chars })
    })
        .then(r => r.json())
        .then(data => {
            const out = document.querySelector('.asciiOutput');
            if (out && data.Output !== undefined) { out.innerHTML = data.Output; }
        })
        .catch(e => console.error(e));
}

// Calculates maximum screen characters, prints to HTML input, and dispatches via fetch ONLY if realtime is active
function sendDynamicMaxCharacters() {
    const outputElement = document.querySelector(".asciiOutput");
    const maxCharsInput = document.getElementById('max-chars-input');
    if (!outputElement) { return; }

    const computedStyle = window.getComputedStyle(outputElement);
    const fontFamily = computedStyle.fontFamily;
    const fontSize = computedStyle.fontSize;
    const fontWeight = computedStyle.fontWeight;
    const availableWidth = window.innerWidth;

    const canvas = document.createElement("canvas");
    const context = canvas.getContext("2d");
    if (context) {
        context.font = `${fontWeight} ${fontSize} ${fontFamily}`;
        const characterWidth = context.measureText("W").width;
        const maxCharacters = Math.floor(availableWidth / characterWidth);

        if (maxCharsInput) {
            maxCharsInput.value = maxCharacters;
        }

        // Run state synchronization to capture fresh settings layout values
        syncState();

        // Strict block condition: If realtime is disabled, prevent computing and sending max characters
        if (!STATE.realtime) { return; }

        const { font_wrap, font_align, active_font, realtime, user_text } = STATE;
        fetch("/api/session-state", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ font_wrap, font_align, active_font, realtime, user_text, max_chars: maxCharacters }),
        })
            .then(response => response.json())
            .then(data => {
                const out = document.querySelector('.asciiOutput');
                if (out && data.output !== undefined) { out.textContent = data.output; }
            })
            .catch(error => console.error("Error sending characters to Go:", error));
    }
}

// Minimal initialization and Event Mapping
let debounce;
let resizeDebounceTimer;
let lastWidth = window.innerWidth;

window.addEventListener('load', () => {
    const textarea = document.getElementById('user-text');

    // Run explicit evaluation tracking right on loading cycles
    syncState();
    sendDynamicMaxCharacters();

    // Settings changes: These update state AND trigger a fetch request ONLY if realtime is currently active
    const elements = document.querySelectorAll('#font-wrap-toggle, #font-align-select, input[name="font"]');
    elements.forEach(el => {
        el.addEventListener('change', () => {
            syncState();
            sendUpdate();
        });
    });

    // FIXED: Realtime toggle event tracking mapped strictly to the correct name attribute filter
    const realtimeToggle = document.querySelector('input[name="Realtime"]') || document.getElementById('realtime-toggle');
    if (realtimeToggle) {
        realtimeToggle.addEventListener('change', () => {
            syncState();
            // Only fire update fetch if it was just turned ON
            if (STATE.realtime) {
                sendUpdate();
            }
        });
    }

    if (textarea) {
        textarea.addEventListener('input', () => {
            clearTimeout(debounce);
            debounce = setTimeout(sendUpdate, 250);
        });
    }

    window.addEventListener("resize", () => {
        const currentWidth = window.innerWidth;

        // Trigger only if the horizontal width actually changed (ignoring height changes from collapses)
        if (currentWidth !== lastWidth) {
            lastWidth = currentWidth; // Update the cached width immediately

            clearTimeout(resizeDebounceTimer);
            resizeDebounceTimer = setTimeout(() => {
                console.log("[Resize] Width changed to " + currentWidth + "px. Re-calculating characters.");
                sendDynamicMaxCharacters();
            }, 400);
        }
    });
});


// Completely wipes frontend targets, updates state, and submits a clean traditional POST request to wipe Go memory
const resetBtn = document.getElementById('reset-btn');
if (resetBtn) {
    resetBtn.addEventListener('click', (e) => {
        e.preventDefault(); // Stop any default button actions

        const textarea = document.getElementById('user-text');
        const form = document.querySelector('footer form');

        // 1. Clear frontend textarea immediately
        if (textarea) {
            textarea.value = "";
        }

        // 2. Clear javascript global tracking states
        STATE.user_text = "";
        STATE.realtime = false; // Turn off realtime temporary to prevent race-condition fetches
        STATE.font_wrap = 'on';
        STATE.font_align = 'left';
        STATE.active_font = 'standard';

        // 3. Force a standard clean Form Post back to Go to wipe its template persistence variables
        if (form) {
            form.submit();
        }
    });
}
