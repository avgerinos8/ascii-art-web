// ==========================================================================
// COLOR PREVIEW & DYNAMIC INDEXING LOGIC
// ==========================================================================

const colorsList = document.getElementById('colors-list');
const addColorBtn = document.getElementById('add-color-btn');

// Helper: Converts HSL values into a HEX string format
function hslToHex(h, s, l) {
    l /= 100;
    const a = (s * Math.min(l, 1 - l)) / 100;
    const f = n => {
        const k = (n + h / 30) % 12;
        const color = l - a * Math.max(Math.min(k - 3, 9 - k, 1), -1);
        return Math.round(255 * color).toString(16).padStart(2, '0');
    };
    return `#${f(0)}${f(8)}${f(4)}`;
}

// Main logic updated to fix overlapping z-indexing and apply exact slider center colors
function updateColorForm(formElement) {
    const h = parseInt(formElement.querySelector('.c-hue').value);
    const s = parseInt(formElement.querySelector('.c-sat').value);
    const l = parseInt(formElement.querySelector('.c-light').value);

    // Dynamic color injection for the selector handles
    formElement.style.setProperty('--thumb-hue', `hsl(${h}, 100%, 50%)`);
    formElement.style.setProperty('--thumb-sat', `hsl(${h}, ${s}%, 50%)`);
    formElement.style.setProperty('--thumb-light', `hsl(${h}, ${s}%, ${l}%)`);

    // CREATIVE FREEDOM: Interactive realistic track fluid gradients mapping
    // Saturation slider track shows the transition from pure gray to fully saturated color at current Lightness
    formElement.style.setProperty('--sat-gradient', `linear-gradient(to right, hsl(${h}, 0%, ${l}%) 0%, hsl(${h}, 100%, ${l}%) 100%)`);

    // Lightness slider track shows full spectrum transition from pure black, through current active Saturation color, to pure white
    formElement.style.setProperty('--light-gradient', `linear-gradient(to right, hsl(${h}, ${s}%, 0%) 0%, hsl(${h}, ${s}%, 50%) 50%, hsl(${h}, ${s}%, 100%) 100%)`);

    // Sync visual response layers
    const previewBox = formElement.querySelector('.preview-box');
    if (previewBox) {
        previewBox.style.backgroundColor = `hsl(${h}, ${s}%, ${l}%)`;
    }

    const hiddenHex = formElement.querySelector('.hidden-hex');
    if (hiddenHex) {
        hiddenHex.value = hslToHex(h, s, l);
    }
}

// Updates text area placeholders sequentially to look clean at a massive 32px font scale mapping
function updatePlaceholders() {
    const inputs = colorsList.querySelectorAll('.substringInput');
    inputs.forEach((input, index) => {
        input.placeholder = `Color ${index + 1}`;
    });
}

// Handles the checkbox "All" toggle behaviors
function handleAllCheckboxToggle(checkbox) {
    const formElement = checkbox.closest('.colorForm');
    if (!formElement) return;

    const textarea = formElement.querySelector('.substringInput');
    if (!textarea) return;

    if (checkbox.checked) {
        textarea.dataset.oldValue = textarea.value; // Save current text if any
        textarea.value = "_ALL_TEXT_"; // Special flag for both Go backend form parser and Fetch JSON
        textarea.disabled = false; // MUST be false so it remains accessible to fetch.js and form POST!
        textarea.readOnly = true; // Block manual keyboard typing inputs

        // Hides the textarea visually but keeps its layout space intact
        textarea.style.visibility = "hidden";
    } else {
        textarea.value = textarea.dataset.oldValue || ""; // Restore text
        textarea.readOnly = false;
        textarea.disabled = false;

        // Brings back the textarea to its normal state
        textarea.style.visibility = "visible";
    }
}

// Click listener to inject the absolute element structure
addColorBtn.addEventListener('click', () => {
    const colorDiv = document.createElement('div');
    colorDiv.className = 'colorForm';

    // FIXED: Generated a dynamic random integer between 0 and 360 for Hue channel spectrum
    const randomHue = Math.floor(Math.random() * 361);

    // FIXED: Restored Saturation back to the requested 80% standard value configuration
    const defaultSaturation = 80;

    // FIXED: Calculate initial matching Hex Code dynamically based on the random hue variable
    const initialHex = hslToHex(randomHue, defaultSaturation, 50);

    colorDiv.innerHTML = `
    <div class="colorSubstringInfo">
      <textarea class="substringInput" name="substring[]" rows="1"></textarea>
      <label class="checkbox-container global-outfit-text">
        <input type="checkbox" name="AllText[]" value="on" class="all-text-check" />
        <span>All</span>
      </label>
    </div>

    <div class="sliders-block global-outfit-text">
      <div class="slider-row">
        <!-- FIXED: Set value attribute to the generated randomHue variable context mapping -->
        <input type="range" class="c-hue" min="0" max="360" value="${randomHue}">
      </div>
      <div class="slider-row">
        <!-- FIXED: Reverted back to the solid 80% saturation default property channel -->
        <input type="range" class="c-sat" min="0" max="100" value="${defaultSaturation}">
      </div>
      <div class="slider-row">
        <input type="range" class="c-light" min="0" max="100" value="50">
      </div>
    </div>

    <div class="preview-box"></div>
    <!-- FIXED: Bound value straight into the dynamically calculated initialHex string -->
    <input type="hidden" class="hidden-hex" name="hexcolorcode[]" value="${initialHex}">
    
    <button type="button" class="delete-color-btn">x</button>
  `;

    colorsList.insertBefore(colorDiv, addColorBtn);

    // FIXED: Wrapped update function inside a browser animation frame cycle thread loop context.
    // This strictly forces the application of the custom generated random Hue gradients layout properties instantly.
    requestAnimationFrame(() => {
        updateColorForm(colorDiv);
    });

    updatePlaceholders();
});

// Real-time slide tracking via bubbling delegation
colorsList.addEventListener('input', (e) => {
    if (e.target.type === 'range') {
        const formElement = e.target.closest('.colorForm');
        updateColorForm(formElement);
    }
});

// Change listener tracking for dynamic element checkbox clicks
colorsList.addEventListener('change', (e) => {
    if (e.target.classList.contains('all-text-check')) {
        handleAllCheckboxToggle(e.target);
    }
});

// Click delegation selector fallback logic to handle the custom close actions
colorsList.addEventListener('click', (e) => {
    if (e.target.classList.contains('delete-color-btn')) {
        const wrapper = e.target.closest('.colors-scroll-wrapper');

        e.target.parentElement.remove();
        updatePlaceholders();

        // Trigger a dynamic change event so fetch.js knows a panel was deleted
        if (wrapper) {
            wrapper.dispatchEvent(new Event('change', { bubbles: true }));
        }
    }
});

// Transform vertical mouse wheel movements into fluid horizontal scrolling
colorsList.addEventListener('wheel', (e) => {
    if (e.deltaY !== 0) {
        e.preventDefault(); // Intercept and block default native page body container scroll bounds
        colorsList.scrollLeft += e.deltaY; // Push horizontal displacement axis
    }
});

// Prevent Enter key updates on any substring textareas to lock them to a single line
document.getElementById('colors-list').addEventListener('keydown', (e) => {
    if (e.target.classList.contains('substringInput') && e.key === 'Enter') {
        e.preventDefault(); // Blocks the newline insertion character entirely
    }
});
