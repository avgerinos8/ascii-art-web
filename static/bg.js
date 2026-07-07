// ── Background ASCII Vignette — Breathing + Mouse Entropy ────────⊃

// Helper function to convert HEX color string (6 or 8 characters) to RGB object
function hexToRgb(hex) {
    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})?$/i.exec(hex.trim());
    return result ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16)
    } : { r: 145, g: 145, b: 145 };
}

// Helper function to determine the correct CSS variable based on the active theme
function getActiveThemeColor() {
    const isDark = document.body.getAttribute('data-theme') === 'dark';
    const varName = isDark ? '--dark-bg-effe' : '--light-bg-effe';
    return getComputedStyle(document.documentElement).getPropertyValue(varName).trim();
}

let currentHex = getActiveThemeColor();
let rgbColor = hexToRgb(currentHex);

window.addEventListener('themeChanged', () => {
    setTimeout(() => {
        currentHex = getActiveThemeColor();
        rgbColor = hexToRgb(currentHex);
    }, 0);
});

const CHARS = [' ', ' ', '░', '░', '▒', '▓', '█'];

let BGSIZE = 0.5;
let BgDistance = 2;
let FS = 16 * BGSIZE;
let LINE_HEIGHT = FS * BgDistance;
let CW = FS * 0.6;
const OP = 0.35;
let NOISE = 0.30;

let BREATHE_SPEED = 0.0004;
const WAVE_SPREAD = 0.6;
let BREATHE_DEPTH = 0.25;

const MOUSE_RADIUS = 12;
const MOUSE_PEAK = 0.4;
const MOUSE_DECAY = 0.001;
const SOFTNESS = 7.5;

const canvas = document.createElement('canvas');
canvas.style.cssText = 'position:fixed;top:0;left:0;z-index:0;pointer-events:none;user-select:none;';
document.body.appendChild(canvas);

const ctx = canvas.getContext('2d');

let noiseGrid = [];
let heatGrid = [];

// Track active alignment state reactively
let activeAlign = 'left';

function buildGrid() {
    const cols = Math.floor(canvas.width / CW) + 1;
    const rows = Math.floor(canvas.height / LINE_HEIGHT) + 1;
    noiseGrid = [];
    heatGrid = [];
    for (let y = 0; y < rows; y++) {
        noiseGrid[y] = [];
        heatGrid[y] = [];
        for (let x = 0; x < cols; x++) {
            noiseGrid[y][x] = {
                spatialNoise: (Math.random() - 0.5) * NOISE,
                phaseOffset: Math.random() * Math.PI * 2,
            };
            heatGrid[y][x] = 0;
        }
    }
}

// ── Reactive Vignette Alignment Sync ────────────────────────────────────────⊃
function syncVignetteAlignment(force = false) {
    const realtimeToggle = document.querySelector('input[name="Realtime"]') || document.getElementById('realtime-toggle');
    const isRealtime = realtimeToggle ? realtimeToggle.checked : false;

    // Halt any background alignment shifts if realtime is disabled, unless it's the initial page load
    if (!isRealtime && !force) {
        return;
    }

    const wrapToggle = document.getElementById('font-wrap-toggle');
    const alignSelect = document.getElementById('font-align-select');

    // If wrap is unchecked (off), force left alignment behavior instantly
    if (wrapToggle && !wrapToggle.checked) {
        activeAlign = 'left';
    } else if (alignSelect) {
        activeAlign = alignSelect.value || 'left';
    } else {
        activeAlign = 'left';
    }
}

window.addEventListener('mousemove', e => {
    const cols = heatGrid[0]?.length ?? 0;
    const rows = heatGrid.length;

    const gx = e.clientX / CW;
    const gy = e.clientY / LINE_HEIGHT;

    const igx = Math.floor(gx);
    const igy = Math.floor(gy);

    for (let dy = -MOUSE_RADIUS; dy <= MOUSE_RADIUS; dy++) {
        for (let dx = -MOUSE_RADIUS; dx <= MOUSE_RADIUS; dx++) {
            const nx = igx + dx;
            const ny = igy + dy;
            if (nx < 0 || ny < 0 || nx >= cols || ny >= rows) continue;

            const screenDx = dx * CW;
            const screenDy = dy * LINE_HEIGHT;
            const screenR = MOUSE_RADIUS * Math.max(CW, LINE_HEIGHT);

            const d2 = (screenDx * screenDx + screenDy * screenDy) / (screenR * screenR);
            const heat = Math.exp(-d2 * SOFTNESS) * MOUSE_PEAK;

            if (heat > heatGrid[ny][nx]) heatGrid[ny][nx] = heat;
        }
    }
});

function draw(timestamp) {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    ctx.font = `${FS}px monospace`;

    const cols = Math.floor(canvas.width / CW) + 1;
    const rows = Math.floor(canvas.height / LINE_HEIGHT) + 1;

    // ── Dynamic Vignette Anchor System ──────────────────────────────────────⊃
    let cx = cols / 2;
    const cy = rows / 2;

    if (activeAlign === 'left') {
        cx = 0;        // Subtle on the left edge, strong on the right
    } else if (activeAlign === 'right') {
        cx = cols;     // Subtle on the right edge, strong on the left
    } else {
        cx = cols / 2; // Center / Justify behavior (subtle in the center)
    }

    // Dynamic max distance scaling from the active focal point to any viewport corner
    const maxD = Math.max(
        Math.hypot(0 - cx, 0 - cy),
        Math.hypot(cols - cx, 0 - cy),
        Math.hypot(0 - cx, rows - cy),
        Math.hypot(cols - cx, rows - cy)
    );

    for (let y = 0; y < rows; y++) {
        for (let x = 0; x < cols; x++) {
            const cell = noiseGrid[y]?.[x];
            if (!cell) continue;

            const dist = Math.hypot(x - cx, y - cy) / maxD;
            const wave = Math.sin(timestamp * BREATHE_SPEED + dist * WAVE_SPREAD + cell.phaseOffset);

            const heat = heatGrid[y]?.[x] ?? 0;
            const t = Math.max(0, Math.min(1, 0.75 * dist + cell.spatialNoise + wave * BREATHE_DEPTH + heat));

            const charIndex = Math.floor(t * (CHARS.length - 1));
            const opacity = Math.min(1, OP + wave * 0.08 + heat * 0.2);

            ctx.fillStyle = `rgba(${rgbColor.r}, ${rgbColor.g}, ${rgbColor.b}, ${opacity.toFixed(3)})`;
            ctx.fillText(CHARS[charIndex], x * CW, (y + 1) * LINE_HEIGHT);

            if (heatGrid[y][x] > 0) heatGrid[y][x] = Math.max(0, heatGrid[y][x] - MOUSE_DECAY);
        }
    }

    requestAnimationFrame(draw);
}

// ── Event Mapping & Listeners ───────────────────────────────────────────────⊃
function setupSliderListeners() {
    const sizeInput = document.querySelector('input[name="BgSize"]');
    const distInput = document.querySelector('input[name="BgDistance"]');
    const speedInput = document.querySelector('input[name="BreatheSpeed"]');
    const noiseInput = document.querySelector('input[name="Noise"]');
    const depthInput = document.querySelector('input[name="BreatheDepth"]');

    if (sizeInput) {
        sizeInput.addEventListener('input', (e) => {
            BGSIZE = parseFloat(e.target.value);
            FS = 16 * BGSIZE;
            CW = FS * 0.6;
            LINE_HEIGHT = FS * BgDistance;
            buildGrid();
        });
    }

    if (distInput) {
        distInput.addEventListener('input', (e) => {
            BgDistance = parseFloat(e.target.value);
            LINE_HEIGHT = FS * BgDistance;
            buildGrid();
        });
    }

    if (speedInput) {
        speedInput.addEventListener('input', (e) => {
            BREATHE_SPEED = parseFloat(e.target.value);
        });
    }

    if (noiseInput) {
        noiseInput.addEventListener('input', (e) => {
            NOISE = parseFloat(e.target.value);
            buildGrid();
        });
    }

    if (depthInput) {
        depthInput.addEventListener('input', (e) => {
            BREATHE_DEPTH = parseFloat(e.target.value);
        });
    }
}

// Attach event tracking directly to UI components for instant synchronization
function setupStateListeners() {
    const wrapToggle = document.getElementById('font-wrap-toggle');
    const alignSelect = document.getElementById('font-align-select');
    const realtimeToggle = document.querySelector('input[name="Realtime"]') || document.getElementById('realtime-toggle');

    if (wrapToggle) {
        wrapToggle.addEventListener('change', () => syncVignetteAlignment(false));
    }
    if (alignSelect) {
        alignSelect.addEventListener('change', () => syncVignetteAlignment(false));
    }
    if (realtimeToggle) {
        // Force evaluation if realtime is enabled layout-wide reactively
        realtimeToggle.addEventListener('change', () => syncVignetteAlignment(false));
    }
}

// ── main ────────────────────────────────────────────────────────────────────⊃
function init() {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const sizeInput = document.querySelector('input[name="BgSize"]');
    const distInput = document.querySelector('input[name="BgDistance"]');
    const noiseInput = document.querySelector('input[name="Noise"]');
    const speedInput = document.querySelector('input[name="BreatheSpeed"]');
    const depthInput = document.querySelector('input[name="BreatheDepth"]');

    // Try to load saved settings from persistence storage if DOM inputs are missing (like in error.html)
    let savedState = {};
    try {
        const raw = localStorage.getItem('ascii_web_settings');
        if (raw) savedState = JSON.parse(raw);
    } catch (e) {
        console.error("Failed to parse storage inside bg.js", e);
    }

    // 1st choice: DOM Input | 2nd choice: LocalStorage | 3rd choice: Hardcoded Index HTML Defaults
    BGSIZE = sizeInput ? parseFloat(sizeInput.value) : (savedState.BgSize !== undefined ? parseFloat(savedState.BgSize) : 1.5);
    BgDistance = distInput ? parseFloat(distInput.value) : (savedState.BgDistance !== undefined ? parseFloat(savedState.BgDistance) : 1.0);
    NOISE = noiseInput ? parseFloat(noiseInput.value) : (savedState.Noise !== undefined ? parseFloat(savedState.Noise) : 0.30);
    BREATHE_SPEED = speedInput ? parseFloat(speedInput.value) : (savedState.BreatheSpeed !== undefined ? parseFloat(savedState.BreatheSpeed) : 0.0004);
    BREATHE_DEPTH = depthInput ? parseFloat(depthInput.value) : (savedState.BreatheDepth !== undefined ? parseFloat(savedState.BreatheDepth) : 0.25);

    FS = 16 * BGSIZE;
    LINE_HEIGHT = FS * BgDistance;
    CW = FS * 0.6;

    // Retry theme fetch after a small macro-task delay to handle CSS variable bindings on error layouts
    setTimeout(() => {
        currentHex = getActiveThemeColor();
        rgbColor = hexToRgb(currentHex);
    }, 50);

    syncVignetteAlignment(true);
    buildGrid();
    setupSliderListeners();
    setupStateListeners();
    requestAnimationFrame(draw);
}

window.addEventListener('resize', () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    buildGrid();
});

init();